package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func echo_handler(conn net.Conn) {
	defer conn.Close()
	b := []byte{}
	conn.Read(b)
	fmt.Println(string(b))
	io.WriteString(conn, "response from server")
}

func main() {
	ln, e := net.Listen("tcp", ":23")
	if e != nil {
		log.Fatal(e)
		return
	}
	for {
		conn, e := ln.Accept()
		if e != nil {
			log.Fatal(e)
			return
		}
		go echo_handler(conn)
	}
}
