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
		command, args := input[0], input[1:]
		fmt.Println(input)

		// minimum implementation
		// https://datatracker.ietf.org/doc/html/rfc959#section-5
		switch command {
		case "USER":
			err = fc.user(args[0])
		case "QUIT":
			err = fc.quit()
			// forから抜けるためエラー処理後にgotoする
			if err != nil {
				// TODO: error handling
				fmt.Printf("%v\n", err)
			}
			goto L
		case "PORT":
			fmt.Println(command)
		case "TYPE":
			fmt.Println(command)
		case "MODE":
			fmt.Println(command)
		case "STRU":
			fmt.Println(command)
		case "RETR":
			fmt.Println(command)
		case "STOR":
			fmt.Println(command)
		case "NOOP":
			err = fc.noop()
		default:
			fc.reply(status500)
		}

		if err != nil {
			// TODO: error handling
			fmt.Printf("%v\n", err)
			break
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
