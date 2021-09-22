#!/bin/bash

echo "go run ./mandelbrot.go 0 > out_complex64.png"
go run ./mandelbrot.go 0 > out_complex64.png

echo "go run ./mandelbrot.go 1 > out_complex128.png"
go run ./mandelbrot.go 1 > out_complex128.png

echo "go run ./mandelbrot.go 2 > out_bigfloat.png"
go run ./mandelbrot.go 2 > out_bigfloat.png

echo "go run ./mandelbrot.go 3 > out_bigrat.png"
go run ./mandelbrot.go 3 > out_bigrat.png