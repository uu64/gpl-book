package popcount2

import (
	"testing"

	"github.com/uu64/gpl-book/ch02/ex03/popcount"
)

var input = uint64(13)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(input)
	}
}

func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount2(input)
	}
}

func TestPopCount2(t *testing.T) {
	var tests = []struct {
		input uint64
		want  int
	}{
		{0, 0},
		{18446744073709551615, 64},
		{13, 3},
	}
	for _, test := range tests {
		if got := PopCount2(test.input); got != test.want {
			t.Errorf("PopCount(%d) = %d\n", test.input, got)
		}
	}
}
