package main

import (
	"fmt"

	"github.com/uu64/gpl-book/ch07/ex01/counter"
)

func main() {
	var wc counter.WordCounter
	wc.Write([]byte("hello, world!\nこんにちは\n"))
	fmt.Println(wc)

	var lc counter.LineCounter
	lc.Write([]byte("hello, world!\nこんにちは\n"))
	fmt.Println(lc)
}
