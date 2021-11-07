package main

import (
	"fmt"

	"github.com/uu64/gpl-book/ch07/ex03/treesort"
)

func main() {
	list := []int{2, 15, 7, 3, 9, 10}
	fmt.Println(list)
	treesort.Sort(list)
	fmt.Println(list)
}
