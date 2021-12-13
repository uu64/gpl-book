package ftp

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type Server struct {
	Addr string
}

func (srv *Server) listen() (net.Listener, error) {
	addr := srv.Addr
	if addr == "" {
		addr = ":21"
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	return ln, nil
}

func (srv *Server) serve(ln net.Listener) error {
	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}

		ftpConn, err := newConn(conn)
		if err != nil {
			return err
		}
		go srv.handler(ftpConn)
	}
}

func (srv *Server) handler(fc *ftpConn) {
	defer fc.close()

	// TODO: error handling
	fc.accept()

	s := bufio.NewScanner(fc.conn)
	for s.Scan() {
		var status string
		var err error

		input := strings.Fields(s.Text())
		command, args := strings.ToUpper(input[0]), input[1:]

		fmt.Println(input)

		switch command {
		// NOTE:
		// FILE TRANSFER PROTOCOL (FTP) / minimum implementation
		// https://datatracker.ietf.org/doc/html/rfc959#section-5
		case "USER":
			status, err = fc.user(args)
		case "QUIT":
			status, err = fc.quit()
			// TODO: error handling
			fc.reply(status)
			// forから抜けるためエラー処理後にgotoする
			if err != nil {
				// TODO: error handling
				fmt.Printf("%v\n", err)
			}
			goto L
		case "PORT":
			status, err = fc.port(args)
		case "TYPE":
			status, err = fc.dataType(args)
		case "MODE":
			status, err = fc.mode(args)
		case "STRU":
			status, err = fc.stru(args)
		case "RETR":
			status, err = fc.retr(args)
		case "STOR":
			status, err = fc.stor(args)
		case "NOOP":
			status, err = fc.noop()
		// NOTE:
		// FTP Extensions for IPv6 and NATs
		// https://datatracker.ietf.org/doc/html/rfc2428
		case "EPRT":
			status, err = fc.eprt(args)
		default:
			status = status502
			err = fmt.Errorf("command is not supported: %s", command)
		}

		fmt.Println(status)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		if status != "" {
			// TODO: error handling
			fc.reply(status)
		}
	}
L:
}

func ListenAndServe(addr string) error {
	server := &Server{Addr: addr}
	ln, err := server.listen()
	if err != nil {
		return err
	}
	return server.serve(ln)
}
