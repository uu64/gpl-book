package main

import (
	"fmt"
	"os"

	"github.com/uu64/gpl-book/ch05/ex03/showtext"
	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}

	showtext.ShowText(os.Stdout, doc)
}
