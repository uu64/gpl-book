package main

import "testing"

func TestExpand(t *testing.T) {
	tests := []struct {
		s    string
		f    func(string) string
		want string
	}{
		{"$test", toUpper, "TEST"},
		{"ab$test", toUpper, "abTEST"},
		{"one two $three", toUpper, "one two THREE"},
		{"$Hello, World!", toUpper, "HELLO, World!"},
		{"$12_345?", reverse, "543_21?"},
		{"apple $orange banana", toUpper, "apple ORANGE banana"},
	}
	for _, test := range tests {
		got := expand(test.s, test.f)
		if got != test.want {
			t.Errorf("s: %s, want: %s, got: %s\n", test.s, test.want, got)
		}
	}
}
