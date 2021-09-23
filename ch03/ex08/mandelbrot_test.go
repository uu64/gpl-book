package main

import (
	"testing"

	"github.com/uu64/gpl-book/ch03/ex08/cmplxbigfloat"
	"github.com/uu64/gpl-book/ch03/ex08/cmplxbigrat"
)

const (
	x          = float64(0)/width*(xmax-xmin) + xmin
	y          = float64(0)/height*(ymax-ymin) + ymin
	iterations = 10
)

func BenchmarkMandelbrotCmplx64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		z := complex64(complex(x, y))
		_ = mandelbrotCmplx64(z, iterations)
	}
}

func BenchmarkMandelbrotCmplx128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		z := complex(x, y)
		_ = mandelbrotCmplx128(z, iterations)
	}
}

func BenchmarkMandelbrotBigFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		z := complex(x, y)
		_ = mandelbrotBigFloat(cmplxbigfloat.New(z), iterations)
	}
}

func BenchmarkMandelbrotBigRat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		z := complex(x, y)
		_ = mandelbrotBigRat(cmplxbigrat.New(z), iterations)
	}
}
