package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type ntpServer struct {
	timezone string
	address  string
}

type worldTime struct {
	timezone   string
	timeString string
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Specify the time zone and the server to connect to as arguments.")
		fmt.Println("ex) clockwall NewYork=localhost:8010 Tokyo=localhost:8020 London=Localhost:8030")
		return
	}

	servers := []*ntpServer{}
	for _, arg := range os.Args[1:] {
		if !strings.Contains(arg, "=") {
			log.Fatalf("invalid arguments: %s\n", arg)
		}
		result := strings.Split(arg, "=")

		loc, err := time.LoadLocation(result[0])
		if err != nil {
			log.Fatalf("invalid timezone: %s\n", result[0])
		}
		servers = append(servers, &ntpServer{loc.String(), result[1]})
	}

	ch := make(chan *worldTime)
	for _, s := range servers {
		go handleConn(s, ch)
	}

	print(servers, ch)
}

func handleConn(server *ntpServer, ch chan<- *worldTime) {
	conn, err := net.Dial("tcp", server.address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ch <- &worldTime{
			timezone:   server.timezone,
			timeString: scanner.Text(),
		}
	}
}

func print(servers []*ntpServer, ch <-chan *worldTime) {
	for {
		m := make(map[string]string)
		for i := 0; i < len(servers); i++ {
			wt := <-ch
			m[wt.timezone] = wt.timeString
		}

		// cursor reset
		fmt.Printf("\033[2J")
		fmt.Printf("\033[0;0H")

		w := 20
		for _, s := range servers {
			time := m[s.timezone]
			p := strings.Repeat(" ", (w-len(time))/2)
			// bold font
			fmt.Printf("\033[1m%s%s%s\033[0m", p, time, p)
		}
		fmt.Println()

		for _, s := range servers {
			tz := s.timezone
			p := strings.Repeat(" ", (w-len(tz))/2)
			// italic font
			fmt.Printf("\033[3m%s%s%s\033[0m", p, tz, p)
		}
		fmt.Println()
	}
}
