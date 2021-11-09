package main

import (
	"flag"
	"fmt"

	"github.com/uu64/gpl-book/ch07/ex06/tempconv"
)

type Kelvin float64

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
