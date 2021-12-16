package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	stages, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(runPipeline(stages))
}

func runPipeline(stages int) int {
	var count int
	var prev, next chan int

	start := make(chan int)
	prev = start
	for count < stages {
		next = make(chan int)

		go func(prev <-chan int, next chan<- int) {
			x := <-prev
			next <- x + 1
		}(prev, next)

		prev = next
		count++
	}
	start <- 0
	result := <-prev
	return result
}
