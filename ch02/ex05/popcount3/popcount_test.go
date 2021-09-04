package popcount3

import (
	"math"
	"testing"

	"github.com/uu64/gpl-book/ch02/ex03/popcount"
	"github.com/uu64/gpl-book/ch02/ex04/popcount2"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(13)
	}
}

func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount2.PopCount2(13)
	}
}

func BenchmarkPopCount3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount3(13)
	}
}

func BenchmarkPopCountWithUintMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0)
	}
}

func BenchmarkPopCount2WithUintMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount2.PopCount2(0)
	}
}

func BenchmarkPopCount3WithUintMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount3(0)
	}
}

func BenchmarkPopCountWithUintMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(math.MaxUint64)
	}
}

func BenchmarkPopCount2WithUintMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount2.PopCount2(math.MaxUint64)
	}
}

func BenchmarkPopCount3WithUintMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount3(math.MaxUint64)
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
