package main

import (
	"unicode/utf8"
)

func reverseInner(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverse(s []byte) {
	for i := 0; i < len(s); {
		_, size := utf8.DecodeRune(s[i:])
		reverseInner(s[i : i+size])
		i += size
	}
	reverseInner(s)
}
