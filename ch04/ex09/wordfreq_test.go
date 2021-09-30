package main

import (
	"io"
	"strings"
	"testing"
)

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

func TestWordfreq(t *testing.T) {
	tests := []struct {
		input io.Reader
		want  map[string]int
	}{
		{strings.NewReader(""), make(map[string]int)},
		{strings.NewReader("test"), map[string]int{"test": 1}},
		{strings.NewReader("Hello, world!"), map[string]int{"hello": 1, "world": 1}},
		{strings.NewReader("Eight apes ate eight apples"), map[string]int{"eight": 2, "apes": 1, "ate": 1, "apples": 1}},
	}

	for _, test := range tests {
		if got := wordfreq(test.input); !equal(got, test.want) {
			t.Errorf("wordfreq(%v) = %v\n", test.input, got)
		}
	}
}
