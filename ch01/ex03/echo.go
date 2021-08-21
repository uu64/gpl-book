package main

import (
	"fmt"
	"os"
	"strings"
)

func concat(args []string) string {
	var s, sep string
	for _, v := range args {
		s += sep + v
		sep = " "
	}
	return s
}

func join(args []string) string {
	return strings.Join(args, " ")
}

func main() {
	fmt.Println(concat(os.Args[1:]))
	fmt.Println(join(os.Args[1:]))
}
