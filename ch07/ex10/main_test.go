package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"Hello, World!", false},
		{"level", true},
		{"", true},
		{"aaa", true},
		{"ab", false},
		{"noon", true},
	}
	for _, test := range tests {
		message := sentence([]byte(test.input))
		if got := IsPalindrome(message); got != test.want {
			t.Errorf("IsPalindrome(%s) returns %v\n", test.input, got)
		}
	}
}
