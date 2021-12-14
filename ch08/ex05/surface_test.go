package main

import (
	"io"
	"testing"
)

func BenchmarkDraw(b *testing.B) {
	for i := 0; i < b.N; i++ {
		draw(io.Discard)
	}
}

func BenchmarkConcurrentDraw3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concurrentDraw(io.Discard, 3)
	}
}

func BenchmarkConcurrentDraw5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concurrentDraw(io.Discard, 5)
	}
}

func BenchmarkConcurrentDraw7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concurrentDraw(io.Discard, 7)
	}
}
