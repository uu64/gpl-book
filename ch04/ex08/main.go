package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

const (
	control = "control"
	letter  = "letter"
	mark    = "mark"
	number  = "number"
	punct   = "punct"
	space   = "space"
	symbol  = "symbol"
)

func main() {
	// Unicode Categories
	categories := [...]string{
		control, letter, mark, number, punct, space, symbol,
	}
	counts := make(map[string]int, len(categories)) // counts of Unicode categories
	invalid := 0                                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		switch {
		case unicode.IsControl(r):
			counts[control]++
		case unicode.IsLetter(r):
			counts[letter]++
		case unicode.IsMark(r):
			counts[mark]++
		case unicode.IsNumber(r):
			counts[number]++
		case unicode.IsPunct(r):
			counts[punct]++
		case unicode.IsSpace(r):
			counts[space]++
		case unicode.IsSymbol(r):
			counts[symbol]++
		}
	}
	fmt.Printf("category\tcount\n")
	for _, c := range categories {
		fmt.Printf("%-8q\t%5d\n", c, counts[c])
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
