package ftp

import (
	"fmt"
	"net"
)

type ftpConn struct {
	conn     net.Conn
	username *string
}

func newConn(c net.Conn) *ftpConn {
	return &ftpConn{
		conn:     c,
		username: nil,
	}
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
