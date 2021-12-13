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
		var err error

		input := strings.Fields(s.Text())
		command, args := strings.ToUpper(input[0]), input[1:]

		fmt.Println(input)

		// minimum implementation
		// https://datatracker.ietf.org/doc/html/rfc959#section-5
		switch command {
		case "USER":
			err = fc.user(args)
		case "QUIT":
			err = fc.quit()
			// forから抜けるためエラー処理後にgotoする
			if err != nil {
				// TODO: error handling
				fmt.Printf("%v\n", err)
			}
			goto L
		case "PORT":
			err = fc.setRemotePort(args)
		case "TYPE":
			err = fc.setDataType(args)
		case "MODE":
			err = fc.setTransferMode(args)
		case "STRU":
			err = fc.setDataStructure(args)
		case "RETR":
			fmt.Println(command)
		case "STOR":
			fmt.Println(command)
		case "NOOP":
			err = fc.noop()
		default:
			fc.reply(status502)
			err = fmt.Errorf("command not implemented: %s", command)
		}

		if err != nil {
			fmt.Printf("%v\n", err)
			goto L
		}
	}
L:
	fmt.Println("finish")
}

func ListenAndServe(addr string) error {
	server := &Server{Addr: addr}
	ln, err := server.listen()
	if err != nil {
		return err
	}
	return server.serve(ln)
}
