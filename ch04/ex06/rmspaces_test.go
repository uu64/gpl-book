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

func TestRmdups(t *testing.T) {
	var tests = []struct {
		input  []byte
		result []byte
	}{
		{[]byte("Hello, World!"), []byte("Hello, World!")},
		{[]byte("Hello,　 \tWorld!"), []byte("Hello, World!")},
		{[]byte(" Hello,,World! "), []byte(" Hello,,World! ")},
		{[]byte(" Hello,       ,World! "), []byte(" Hello, ,World! ")},
		{[]byte(" Hello,  	     　,World! "), []byte(" Hello, ,World! ")},
	}

	for _, test := range tests {
		tmp := make([]byte, len(test.input))
		copy(tmp, test.input)
		got := rmspaces(test.input)
		if !equal(test.input[:len(got)], test.result) {
			t.Errorf("rmspaces(%v) = %v\n", string(tmp), string(got))
		}
	}
}
