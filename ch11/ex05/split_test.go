package split_test

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		s   string
		sep string
		len int
	}{
		{"a:b:c", ":", 3},
		{"", ",", 1},
		{"a:b:c", ",", 1},
		{"sample", "", 6},
		{"1,2,3,4", ",", 4},
	}
	for _, test := range tests {
		words := strings.Split(test.s, test.sep)
		if got := len(words); got != test.len {
			t.Errorf("Split(%q, %q) returned %d words, want %d", test.s, test.sep, got, test.len)
		}
	}
}
