package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/uu64/gpl-book/ch02/ex01/tempconv"
)

func conv(s string) {
	t, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
}

func main() {
	if len(os.Args) >= 2 {
		for _, arg := range os.Args[1:] {
			conv(arg)
		}
	} else {
		var s string
		_, err := fmt.Scan(&s)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		conv(s)
	}
}
