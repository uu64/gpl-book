package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

func wordfreq(r io.Reader) map[string]int {
	counts := make(map[string]int)
	input := bufio.NewScanner(r)

	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := strings.TrimFunc(input.Text(), func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsNumber(r)
		})
		counts[strings.ToLower(word)]++
	}
	return counts
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: wordfreq.go FILE_PATH")
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
		os.Exit(1)
	}

	counts := wordfreq(f)
	fmt.Printf("%4v\t%-18v\n", "word", "count")
	for w, c := range counts {
		fmt.Printf("%-4d\t%-18q\n", c, w)
	}
}
