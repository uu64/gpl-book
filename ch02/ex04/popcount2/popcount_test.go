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
