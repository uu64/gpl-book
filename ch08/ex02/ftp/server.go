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
		go srv.handler(conn)
	}
}

func (srv *Server) handler(conn net.Conn) {
	defer conn.Close()
	fmt.Fprintf(conn, "%s", Status220)

	s := bufio.NewScanner(conn)
	for s.Scan() {
		fmt.Println(s.Text())
		input := strings.Fields(s.Text())
		command, _ := input[0], input[1:]
		switch command {
		case "USER":
			fmt.Println("USER")
		default:
			fmt.Println("unknown")
		}
	}
}

func ListenAndServe(addr string) error {
	server := &Server{Addr: addr}
	ln, err := server.listen()
	if err != nil {
		return err
	}
	return server.serve(ln)
}
