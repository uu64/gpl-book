package main

import "testing"

func BenchmarkRunPipeline10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = runPipeline(10000)
	}
}

func BenchmarkRunPipeline100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = runPipeline(100000)
	}
}

func BenchmarkRunPipeline1000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = runPipeline(1000000)
	}
}

// 数分かかっても終わらない
// func BenchmarkRunPipeline10000000(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		_ = runPipeline(10000000)
// 	}
// }
