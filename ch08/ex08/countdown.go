package main

// NOTE: the ticker goroutine never terminates if the launch is aborted.
// This is a "goroutine leak".

import (
	"fmt"
	"os"
	"time"
)

//!+

func main2() {
	// ...create abort channel...

	//!-

	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()

	//!+
	fmt.Println("Commencing countdown.  Press return to abort.")
	select {
	case <-time.After(10 * time.Second):
		// Do nothing.
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}
	// fmt.Println("Commencing countdown.  Press return to abort.")
	// tick := time.Tick(1 * time.Second)
	// for countdown := 10; countdown > 0; countdown-- {
	// 	fmt.Println(countdown)
	// 	select {
	// 	case <-tick:
	// 		// Do nothing.
	// 	case <-abort:
	// 		fmt.Println("Launch aborted!")
	// 		return
	// 	}
	// }
	launch()
}

//!-

func launch() {
	fmt.Println("Lift off!")
}
