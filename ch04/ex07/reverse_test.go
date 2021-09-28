package main

import (
	"fmt"
	"testing"
)

func equal(x, y []byte) bool {
	if len(x) != len(y) {
		fmt.Printf("len %d, %d\n", len(x), len(y))
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			fmt.Printf("v %v, %v\n", x, y)
			return false
		}
	}
	return true
}

func TestReverse(t *testing.T) {
	var tests = []struct {
		input  []byte
		result []byte
	}{
		{[]byte("Hello, World!"), []byte("!dlroW ,olleH")},
		{[]byte("おはよう"), []byte("うよはお")},
		{[]byte("練習問題 4.7"), []byte("7.4 題問習練")},
	}

	for _, test := range tests {
		tmp := make([]byte, len(test.input))
		copy(tmp, test.input)
		reverse(test.input)
		if !equal(test.input, test.result) {
			t.Errorf("reverse(%v), then %v\n", string(tmp), string(test.input))
		}
	}
}
