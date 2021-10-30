package main

import "fmt"

func returnNotZero() (result int) {
	defer func() {
		recover()
		result = 1
	}()
	panic(1)
}

func main() {
	fmt.Println(returnNotZero())
}
