package main

import (
	"io"
	"strings"
	"testing"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func TestLimitReader(t *testing.T) {
	tests := []struct {
		input string
		limit int64
		want  string
		err   error
	}{
		{"Hello, World!", 4, "Hell", nil},
		{"Hello, World!", 0, "", io.EOF},
		{"test", 5, "test", nil},
	}
	for _, test := range tests {
		r := strings.NewReader(test.input)
		lr := LimitReader(r, test.limit)

		tmp := make([]byte, 8)
		n, err := lr.Read(tmp)
		if n != min(int(test.limit), len(test.input)) || err != test.err || strings.Compare(string(tmp[:n]), test.want) != 0 {
			t.Errorf("n = %v err = %v b = %v str = \"%s\"\n", n, err, tmp, string(tmp))
		}
	}
}
