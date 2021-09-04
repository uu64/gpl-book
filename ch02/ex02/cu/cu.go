package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/uu64/gpl-book/ch02/ex01/tempconv"
	"github.com/uu64/gpl-book/ch02/ex02/cu/conv"
)

func cu(s string) {
	t, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cu: %v\n", err)
		os.Exit(1)
	}

	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	m := conv.Meter(t)
	ft := conv.Feet(t)
	kg := conv.Kilogram(t)
	lb := conv.Pound(t)
	format := "%s = %s, %s = %s\n"

	fmt.Printf(format, f, tempconv.FToC(f), c, tempconv.CToF(c))
	fmt.Printf(format, m, conv.MToF(m), ft, conv.FToM(ft))
	fmt.Printf(format, kg, conv.KToP(kg), lb, conv.PToK(lb))
	fmt.Println()
}

func main() {
	if len(os.Args) >= 2 {
		for _, arg := range os.Args[1:] {
			cu(arg)
		}
	} else {
		var s string
		_, err := fmt.Scan(&s)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cu: %v\n", err)
			os.Exit(1)
		}
		cu(s)
	}
}
