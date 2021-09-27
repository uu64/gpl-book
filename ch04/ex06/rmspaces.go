package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func rmspaces(s []byte) []byte {
	prev, size := utf8.DecodeRune(s)
	pos := size
	for i := size; i < len(s); {
		r, size := utf8.DecodeRune(s[i:])
		if !unicode.IsSpace(r) {
			copy(s[pos:], s[i:i+size])
			pos += size
		} else if unicode.IsSpace(r) && !unicode.IsSpace(prev) {
			asciiSpace := []byte(" ")
			copy(s[pos:], asciiSpace)
			pos += len(asciiSpace)
		}
		prev = r
		i += size
	}
	return s[:pos]
}

func main() {
	s := "Hello, 　\tWorld!"
	fmt.Println(string(rmspaces([]byte(s))))
}
