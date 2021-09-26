package main

import "testing"

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func TestRmdups(t *testing.T) {
	var tests = []struct {
		input  []string
		result []string
	}{
		{[]string{"hello", "world"}, []string{"hello", "world"}},
		{[]string{"hello", "hello", "a", "b", "b", "b", "!"}, []string{"hello", "a", "b", "!"}},
	}

	for _, test := range tests {
		if got := rmdups(test.input); !equal(got, test.result) {
			t.Errorf("rmdups(%v) = %v\n", test.input, got)
		}
	}
}
