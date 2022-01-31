package main

import "github.com/uu64/gpl-book/ch12/ex01/display"

func main() {
	m := map[string]int{
		"test":  1,
		"apple": 10,
		"hello": -34,
	}
	display.Display("m", m)

	array := [3]int{
		2, 3, 4,
	}
	display.Display("array", array)

	arraymap := map[[3]int]string{
		array: "hahaha",
	}
	display.Display("arraymap", arraymap)
}
