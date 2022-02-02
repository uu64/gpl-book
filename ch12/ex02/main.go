package main

import (
	"fmt"

	"github.com/uu64/gpl-book/ch12/ex02/display"
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

	type Movie struct {
		Title string
		Year  int
		Ref   [2]int
	}
	movie1 := Movie{"movie1", 2012, [2]int{1, 2}}
	display.Display("movie1", movie1)
	fmt.Println()

	movie2 := Movie{"movie2", 2014, [2]int{4, 2}}
	movie3 := Movie{"movie3", 2009, [2]int{5, 9}}
	structmap := map[Movie]string{
		movie1: "hahaha",
		movie2: "hello",
		movie3: "world",
	}
	display.Display("structmap", structmap)
	fmt.Println()

	type Cyclic struct {
		Name string
		Ref  *Cyclic
	}
	cyclic := Cyclic{"cyclic1", nil}
	cyclic.Ref = &cyclic
	display.Display("cyclic", cyclic)
	fmt.Println()
}
