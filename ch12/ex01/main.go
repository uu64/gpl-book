package main

import (
	"fmt"

	"github.com/uu64/gpl-book/ch12/ex01/display"
)

func main() {
	m := map[string]int{
		"test":  1,
		"apple": 10,
		"hello": -34,
	}
	display.Display("m", m)
	fmt.Println()

	array1 := [3]int{
		2, 3, 4,
	}
	display.Display("array", array1)
	fmt.Println()

	array2 := [3]int{
		4, 5, 6,
	}
	array3 := [3]int{
		7, 8, 9,
	}
	arraymap := map[[3]int]string{
		array1: "hahaha",
		array2: "hello",
		array3: "world",
	}
	display.Display("arraymap", arraymap)
	fmt.Println()
}
