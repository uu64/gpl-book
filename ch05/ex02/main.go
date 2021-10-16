package main

import (
	"fmt"
	"os"

	"github.com/uu64/gpl-book/ch05/ex02/countelem"
	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}

	counter := make(map[string]int)
	countelem.CountElem(counter, doc)
	for k, v := range counter {
		fmt.Printf("%8s: %3d\n", k, v)
	}
}
