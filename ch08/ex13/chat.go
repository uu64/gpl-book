package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

type user struct {
	ch   client
	name string
}

//!+broadcaster
type client chan<- string // an outgoing message channel

var (
	entering = make(chan user)
	leaving  = make(chan user)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]string) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			clients[cli.ch] = cli.name
			members := []string{}
			for _, v := range clients {
				// NOTE: ignoring errors
				members = append(members, v)
			}
			cli.ch <- fmt.Sprintf("All members: %s\n", strings.Join(members, ", "))

		case cli := <-leaving:
			delete(clients, cli.ch)
			close(cli.ch)
		}
	}
}

//!-broadcaster

//!+handleConn
func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- user{ch, who}

	pipe := make(chan string)
	go func() {
		defer log.Printf("close: %s\n", who)
		input := bufio.NewScanner(conn)
		for input.Scan() {
			pipe <- who + ": " + input.Text()
		}
		// NOTE: ignoring potential errors from input.Err()
	}()

L:
	for {
		select {
		case <-time.After(5 * time.Minute):
			break L
		case msg := <-pipe:
			messages <- msg
		}
	}
	leaving <- user{ch, who}
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

//!+main
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
