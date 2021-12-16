package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ping, pong := make(chan struct{}), make(chan struct{})
	done := make(chan struct{})

	var wg sync.WaitGroup
	player := func(send chan<- struct{}, recv <-chan struct{}, done <-chan struct{}) {
		count := 0
		defer func() {
			fmt.Printf("count: %d\n", count)
			wg.Done()
		}()
		for {
			select {
			case <-recv:
				send <- struct{}{}
			case <-done:
				return
			}
			count++
		}
	}

	wg.Add(1)
	go player(ping, pong, done)

	wg.Add(1)
	go player(pong, ping, done)

	ping <- struct{}{}
	time.Sleep(1 * time.Second)
	close(done)

	// 後に閉じるチャネルの通信を受け取る
	select {
	case <-ping:
	case <-pong:
	}

	wg.Wait()
	fmt.Println("finish")
}
