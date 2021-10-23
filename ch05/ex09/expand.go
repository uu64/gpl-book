package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

var pat = regexp.MustCompile(`\$[a-zA-Z0-9]+`)

func findNotWord(b []byte) int {
	for i := 0; i < len(b); {
		r, size := utf8.DecodeRune(b[i:])
		if !unicode.IsNumber(r) && !unicode.IsLetter(r) && byte(r) != '_' {
			return i
		}
		i += size
	}
	return -1
}

func expand(s string, f func(string) string) string {
	b := []byte(s)
	var builder strings.Builder
	for i := 0; i < len(b); {
		r, size := utf8.DecodeRune(b[i:])
		if byte(r) == '$' {
			i += size
			idx := findNotWord(b[i:])
			if idx == -1 {
				idx = len(b) - i
			}
			builder.WriteString(f(string(b[i : i+idx])))
			i += idx
		} else {
			builder.WriteRune(r)
			i += size
		}
	}
	return builder.String()
}

func toUpper(s string) string {
	return strings.ToUpper(s)
}

func reverseInner(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}

func reverse(s string) string {
	b := []byte(s)
	for i := 0; i < len(s); {
		_, size := utf8.DecodeRune(b[i:])
		reverseInner(b[i : i+size])
		i += size
	}
	reverseInner(b)
	return string(b)
}

func main() {
	var s string
	s = "$Hello world!"
	fmt.Println(s)
	fmt.Println(expand(s, toUpper))
	fmt.Println(expand(s, reverse))
}
