#!/bin/bash

echo "go run ./mandelbrot.go 0 > out0.png"
go run ./mandelbrot.go 0 > out0.png

echo "go run ./mandelbrot.go 1 > out1.png"
go run ./mandelbrot.go 1 > out1.png

echo "go run ./mandelbrot.go 2 > out2.png"
go run ./mandelbrot.go 2 > out2.png