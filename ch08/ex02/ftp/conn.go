package ftp

import (
	"fmt"
	"net"
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
}

type transferParameter struct {
	dataType  string
	mode      string
	structure string
}

func newConn(c net.Conn) (*ftpConn, error) {
	_, port, err := net.SplitHostPort(c.RemoteAddr().String())
	if err != nil {
		return nil, err
	}
	return &ftpConn{
		conn: c,
		remote: &remote{
			username: nil,
			port:     &port,
		},
		param: &transferParameter{
			dataType:  typeAscii,
			mode:      modeStream,
			structure: struFile,
		},
	}, nil
}

func (fc *ftpConn) isLogin() bool {
	return fc.remote.username != nil
}

func (fc *ftpConn) reply(status string) error {
	_, err := fmt.Fprintf(fc.conn, "%s\n", status)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (fc *ftpConn) accept() error {
	return fc.reply(status220)
}

func (fc *ftpConn) close() {
	if err := fc.conn.Close(); err != nil {
		// TODO: error handling
		fmt.Println("faild: close")
	}
}

func (fc *ftpConn) user(args []string) error {
	if len(args) != 1 {
		return fc.reply(status501)
	}

	name := args[0]
	fc.remote.username = &name
	// TODO: auth
	return fc.reply(status230)
}

func (fc *ftpConn) quit() error {
	fc.remote.username = nil
	// TODO: データ転送中ならコネクションはまだ閉じない
	// TODO: コントロール接続に関してやることはあるか
	return fc.reply(status221)
}

func (fc *ftpConn) port(args []string) error {
	if !fc.isLogin() {
		return fc.reply(status530)
	}

	if len(args) != 6 {
		return fc.reply(status501)
	}

	h1, h2, h3, h4, p1, p2 := args[0], args[1], args[2], args[3], args[4], args[5]
	addr := fmt.Sprintf("%s.%s.%s.%s", h1, h2, h3, h4)
	fc.remote.addr = &addr

	p1i, err := strconv.Atoi(p1)
	if err != nil {
		return fc.reply(status501)
	}

	p2i, err := strconv.Atoi(p2)
	if err != nil {
		return fc.reply(status501)
	}

	port := strconv.Itoa((p1i * 256) + p2i)
	fc.remote.port = &port

	return fc.reply(status200)
}

func (fc *ftpConn) setType(args []string) error {
	if !fc.isLogin() {
		return fc.reply(status530)
	}

	if len(args) != 1 {
		return fc.reply(status501)
	}

	dataType := args[0]
	if dataType != typeAscii {
		return fc.reply(status504)
	}

	fc.param.dataType = dataType
	return fc.reply(status200)
}

func (fc *ftpConn) noop() error {
	// TODO: データ転送中ならコネクションはまだ閉じない
	// TODO: コントロール接続に関してやることはあるか
	return fc.reply(status200)
}
