package main

import "testing"

func TestGet(t *testing.T) {
	var tests = []struct {
		input string
	}{
		{"http://gopl.io"},
		{"gopl.io"},
	}

	for _, test := range tests {
		err := get(test.input)

		if err != nil {
			t.Errorf("get(%q) %w", test.input, err)
		}
	}
}
