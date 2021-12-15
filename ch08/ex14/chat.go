package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

type user struct {
	ch   client
	name string
}

var clients = make(map[client]string) // all connected clients
var mu sync.Mutex

//!+broadcaster
type client chan<- string // an outgoing message channel

var (
	entering = make(chan user)
	leaving  = make(chan user)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			mu.Lock()
			for cli := range clients {
				cli <- msg
			}
			mu.Unlock()

		case cli := <-entering:
			mu.Lock()
			fmt.Println("entering")
			clients[cli.ch] = cli.name
			members := []string{}
			for _, v := range clients {
				// NOTE: ignoring errors
				members = append(members, v)
			}
			mu.Unlock()
			cli.ch <- fmt.Sprintf("All members: %s\n", strings.Join(members, ", "))

		case cli := <-leaving:
			mu.Lock()
			delete(clients, cli.ch)
			mu.Unlock()
			close(cli.ch)
		}
	}
}

//!-broadcaster

//!+handleConn
func handleConn(conn net.Conn) {
	defer conn.Close()

	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	// Get user name
	ch <- "What is your username?"
	var who string
	input := bufio.NewScanner(conn)
	for input.Scan() {
		isOk := true
		mu.Lock()
		for _, v := range clients {
			if v == input.Text() {
				ch <- fmt.Sprintf("%s is already in use. Specify a other username.", v)
				isOk = false
				break
			}
		}
		mu.Unlock()
		if isOk {
			who = input.Text()
			break
		}
	}

	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- user{ch, who}

	defer func() {
		leaving <- user{ch, who}
		messages <- who + " has left"
		log.Printf("finish: %s\n", who)
	}()

	pipe := make(chan string)
	closed := make(chan struct{})
	go func() {
		defer func() {
			close(closed)
			log.Printf("close: %s\n", who)
		}()
		input := bufio.NewScanner(conn)
		for input.Scan() {
			pipe <- who + ": " + input.Text()
		}
		// NOTE: ignoring potential errors from input.Err()
	}()

	for {
		select {
		case <-closed:
			goto L
		case <-time.After(5 * time.Minute):
			goto L
		case msg := <-pipe:
			messages <- msg
		}
	}
L:
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
