package main

import (
	"image"
	"image/color"
	"image/png"
	"math/big"
	"math/cmplx"
	"os"

	"github.com/uu64/gpl-book/ch03/ex08/cmplxBigFloat"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -0.5555780, -0.5555780, -0.5555750, -0.5555750
		// xmin, ymin, xmax, ymax = -0.74992250820790, 0.02417781581040, -0.74992250820770, 0.02417781581060
		// xmin, ymin, xmax, ymax = -1.0, -1.0, 1.0, 1.0
		width, height = 512, 512
	)
	opt := os.Args[1]
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			switch opt {
			case "0":
				z := complex64(complex(x, y))
				img.Set(px, py, mandelbrotCmplx64(z))
			case "1":
				z := complex(x, y)
				img.Set(px, py, mandelbrotCmplx128(z))
			case "2":
				z := complex(x, y)
				img.Set(px, py, mandelbrotBigFloat(cmplxBigFloat.New(z)))
				// case "3":
				// 	img.Set(px, py, mandelbrot(z))
			}
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrotCmplx64(z complex64) color.Color {
	const iterations = 255
	const contrast = 15

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(complex128(v)) > 2 {
			return color.YCbCr{contrast * n, 255 - contrast*n, contrast * n}
		}
	}
	return color.Black
}

func mandelbrotCmplx128(z complex128) color.Color {
	const iterations = 255
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.YCbCr{contrast * n, 255 - contrast*n, contrast * n}
		}
	}
	return color.Black
}

func mandelbrotBigFloat(z *cmplxBigFloat.Cmplx) color.Color {
	const iterations = 255
	const contrast = 15

	v := cmplxBigFloat.New(complex(0, 0))
	for n := uint8(0); n < iterations; n++ {
		v = cmplxBigFloat.Add(cmplxBigFloat.Mul(v, v), z)
		if cmplxBigFloat.Abs(v).Cmp(big.NewFloat(2)) == 1 {
			return color.YCbCr{contrast * n, 255 - contrast*n, contrast * n}
		}
	}
	return color.Black
}
