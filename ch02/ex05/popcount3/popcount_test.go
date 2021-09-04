package popcount3

import (
	"math"
	"testing"

	"github.com/uu64/gpl-book/ch02/ex03/popcount"
	"github.com/uu64/gpl-book/ch02/ex04/popcount2"
)

var input = uint64(13)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(input)
	}
}

func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount2.PopCount2(input)
	}
}

func BenchmarkPopCount3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount3(input)
	}
}

func TestPopCount3(t *testing.T) {
	var tests = []struct {
		input uint64
		want  int
	}{
		{0, 0},
		{math.MaxUint64, 64},
		{13, 3},
	}
	for _, test := range tests {
		if got := PopCount3(test.input); got != test.want {
			t.Errorf("PopCount(%d) = %d\n", test.input, got)
		}
	}
}
