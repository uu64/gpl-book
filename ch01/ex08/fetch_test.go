package main

import "testing"

func TestGet(t *testing.T) {
	var tests = []struct {
		input string
		want  error
	}{
		{"http://gopl.io", nil},
		{"gopl.io", nil},
	}

	for _, test := range tests {
		if got := get(test.input); got != test.want {
			t.Errorf("get(%q) %v", test.input, got)
		}
	}
}
