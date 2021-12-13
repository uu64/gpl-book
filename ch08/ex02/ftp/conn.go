package ftp

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
)

type ftpConn struct {
	conn   net.Conn
	remote *remote
	param  *transferParameter
}

type remote struct {
	username *string
	addr     *string
	port     *string
	protocol string
}

type transferParameter struct {
	dataType  string
	mode      string
	structure string
}

func newConn(c net.Conn) (*ftpConn, error) {
	host, port, err := net.SplitHostPort(c.RemoteAddr().String())
	if err != nil {
		return nil, err
	}
	return &ftpConn{
		conn: c,
		remote: &remote{
			username: nil,
			addr:     &host,
			port:     &port,
			protocol: protocolIp4,
		},
		param: &transferParameter{
			dataType:  typeAscii,
			mode:      modeStream,
			structure: struFile,
		},
	}, nil
}

func (fc *ftpConn) openDataConn() (net.Conn, error) {
	var network, addr string
	if fc.remote.protocol == protocolIp4 {
		network = "tcp4"
		addr = fmt.Sprintf("%s:%s", *fc.remote.addr, *fc.remote.port)
	} else {
		network = "tcp6"
		addr = fmt.Sprintf("[%s]:%s", *fc.remote.addr, *fc.remote.port)
	}

	return net.Dial(network, addr)
}

func (fc *ftpConn) isLogin() bool {
	return fc.remote.username != nil
}

// TODO: たまに返信に失敗する
func (fc *ftpConn) reply(status string) error {
	_, err := fmt.Fprintf(fc.conn, "%s\n", status)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (fc *ftpConn) close() {
	if err := fc.conn.Close(); err != nil {
		// TODO: error handling
		fmt.Println("faild: close")
	}
}

func (fc *ftpConn) accept() error {
	return fc.reply(status220)
}

func (fc *ftpConn) user(args []string) (status string, err error) {
	if len(args) != 1 {
		status = status501
		return
	}

	name := args[0]
	fc.remote.username = &name
	// TODO: auth
	status = status230
	return
}

func (fc *ftpConn) quit() (status string, err error) {
	fc.remote.username = nil
	// TODO: データ転送中ならコネクションはまだ閉じない
	// TODO: コントロール接続に関してやることはあるか
	status = status221
	return
}

func (fc *ftpConn) noop() (status string, err error) {
	status = status200
	return
}

func (fc *ftpConn) retr(args []string) (status string, err error) {
	if !fc.isLogin() {
		status = status530
		return
	}

	if len(args) != 1 {
		status = status501
		return
	}

	path := args[0]
	info, err := os.Stat(path)
	if err != nil {
		status = status550
		return
	}
	if info.IsDir() {
		status = status550
		err = fmt.Errorf("directory is not supported")
		return
	}

	if err = fc.reply(status150); err != nil {
		return
	}

	conn, err := fc.openDataConn()
	if err != nil {
		status = status425
		return
	}

	defer func() {
		// TODO: error handling
		if err == nil {
			fc.reply(status226)
			status = status250
		} else {
			status = status426
		}
		conn.Close()
	}()

	f, err := os.Open(path)
	if err != nil {
		status = status550
		return
	}

	defer func() {
		err = f.Close()
		if err != nil {
			status = status552
		}
	}()

	// TODO: asciiモードでは改行コードをCRLFにする
	_, err = io.Copy(conn, f)
	if err != nil {
		status = status451
		return
	}

	return
}

func (fc *ftpConn) stor(args []string) (status string, err error) {
	if !fc.isLogin() {
		status = status530
		return
	}

	if len(args) != 1 {
		status = status501
		return
	}

	path := args[0]
	info, err := os.Stat(path)
	if err == nil && info.IsDir() {
		status = status550
		err = fmt.Errorf("target is a directory")
		return
	}

	if err = fc.reply(status150); err != nil {
		return
	}

	conn, err := fc.openDataConn()
	if err != nil {
		status = status425
		return
	}

	defer func() {
		// TODO: error handling
		if err == nil {
			fc.reply(status226)
		} else {
			fc.reply(status426)
		}
		conn.Close()
		fmt.Println("closed")
		status = status250
	}()

	f, err := os.Create(path)
	if err != nil {
		status = status550
		return
	}

	defer func() {
		err = f.Close()
		if err != nil {
			status = status552
		}
	}()

	// NOTE: 複数回同じファイルに書き込むと失敗する
	// TODO: asciiモードでは改行コードをCRLFにする
	_, err = io.Copy(f, conn)
	if err != nil {
		status = status451
		return
	}

	status = status226
	return
}

func (fc *ftpConn) port(args []string) (status string, err error) {
	if !fc.isLogin() {
		status = status530
		return
	}

	if len(args) != 6 {
		status = status501
		return
	}

	h1, h2, h3, h4, p1, p2 := args[0], args[1], args[2], args[3], args[4], args[5]
	addr := fmt.Sprintf("%s.%s.%s.%s", h1, h2, h3, h4)
	fc.remote.addr = &addr

	p1i, err := strconv.Atoi(p1)
	if err != nil {
		status = status501
		return
	}

	p2i, err := strconv.Atoi(p2)
	if err != nil {
		status = status501
		return
	}

	port := strconv.Itoa((p1i * 256) + p2i)
	fc.remote.port = &port

	status = status200
	return
}

func (fc *ftpConn) dataType(args []string) (status string, err error) {
	if !fc.isLogin() {
		status = status530
		return
	}

	if len(args) != 1 {
		status = status501
		return
	}

	dataType := args[0]
	if dataType != typeAscii && dataType != typeImage {
		status = status504
		return
	}

	fc.param.dataType = dataType
	status = status200
	return
}

func (fc *ftpConn) mode(args []string) (status string, err error) {
	if !fc.isLogin() {
		status = status530
		return
	}

	if len(args) != 1 {
		status = status501
		return
	}

	mode := args[0]
	if mode != modeStream {
		status = status504
		return
	}

	fc.param.mode = mode
	status = status200
	return
}

func (fc *ftpConn) stru(args []string) (status string, err error) {
	if !fc.isLogin() {
		status = status530
		return
	}

	if len(args) != 1 {
		status = status501
		return
	}

	stru := args[0]
	if stru != struFile {
		status = status504
		return
	}

	fc.param.structure = stru
	status = status200
	return
}

func (fc *ftpConn) eprt(args []string) (status string, err error) {
	if !fc.isLogin() {
		status = status530
		return
	}

	if len(args) != 1 {
		status = status501
		return
	}

	b := []byte(args[0])
	params := bytes.Split(b[1:len(b)-1], b[0:1])

	if len(params) != 3 {
		status = status501
		return
	}
	protocol := string(params[0])
	addr := string(params[1])
	port := string(params[2])

	if protocol == "1" {
		fc.remote.protocol = protocolIp4
	} else {
		fc.remote.protocol = protocolIp6
	}
	fc.remote.addr = &addr
	fc.remote.port = &port

	status = status200
	return
}
