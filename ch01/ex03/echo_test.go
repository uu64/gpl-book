package main

import "testing"

var words = []string{
	"apple",
	"banana",
	"chocolate",
	"dog",
	"elephant",
}

func BenchmarkConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concat(words)
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		join(words)
	}
}
