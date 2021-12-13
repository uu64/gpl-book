package ftp

import (
	"fmt"
	"net"
)

type ftpConn struct {
	conn   net.Conn
	remote *remote
}

type remote struct {
	username *string
	port     string
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
			port:     port,
		},
	}, nil
}

func (fc *ftpConn) reply(status string) error {
	_, err := fmt.Fprintf(fc.conn, "%s", status)
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

func (fc *ftpConn) user(name string) error {
	fc.remote.username = &name
	// TODO: auth
	return fc.reply(status230)
}

func (fc *ftpConn) quit() error {
	// TODO: データ転送中ならコネクションはまだ閉じない
	// TODO: コントロール接続に関してやることはあるか
	return fc.reply(status221)
}

func (fc *ftpConn) noop() error {
	// TODO: データ転送中ならコネクションはまだ閉じない
	// TODO: コントロール接続に関してやることはあるか
	return fc.reply(status200)
}
