package main

import (
	"image"
	"image/color"
	"image/png"
	"math/big"
	"math/cmplx"
	"os"

	"github.com/uu64/gpl-book/ch03/ex08/cmplxbigfloat"
	"github.com/uu64/gpl-book/ch03/ex08/cmplxbigrat"
)

const (
	// xmin, ymin, xmax, ymax = -0.75100, 0.05100, -0.74600, 0.05600
	// xmin, ymin, xmax, ymax = -0.74996733, 0.05294105, -0.74995971, 0.05294867
	xmin, ymin, xmax, ymax = -0.749966940, 0.052945120, -0.749965420, 0.052946640
	width, height          = 1024, 1024
)

func main() {
	opt := os.Args[1]
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			switch opt {
			case "0":
				z := complex64(complex(x, y))
				img.Set(px, py, mandelbrotCmplx64(z, 255))
			case "1":
				z := complex(x, y)
				img.Set(px, py, mandelbrotCmplx128(z, 255))
			case "2":
				z := complex(x, y)
				img.Set(px, py, mandelbrotBigFloat(cmplxbigfloat.New(z), 255))
			case "3":
				z := complex(x, y)
				img.Set(px, py, mandelbrotBigRat(cmplxbigrat.New(z), 10))
			}
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrotCmplx64(z complex64, iterations uint8) color.Color {
	const contrast = 5

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(complex128(v)) > 2 {
			return color.YCbCr{contrast * n, 255 - contrast*n, contrast * n}
		}
	}
	return color.Black
}

func mandelbrotCmplx128(z complex128, iterations uint8) color.Color {
	const contrast = 5

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.YCbCr{contrast * n, 255 - contrast*n, contrast * n}
		}
	}
	return color.Black
}

func mandelbrotBigFloat(z *cmplxbigfloat.Cmplx, iterations uint8) color.Color {
	const contrast = 5

	v := cmplxbigfloat.New(complex(0, 0))
	for n := uint8(0); n < iterations; n++ {
		v = cmplxbigfloat.Add(cmplxbigfloat.Mul(v, v), z)
		if cmplxbigfloat.Abs(v).Cmp(big.NewFloat(2)) == 1 {
			return color.YCbCr{contrast * n, 255 - contrast*n, contrast * n}
		}
	}
	return color.Black
}

func mandelbrotBigRat(z *cmplxbigrat.Cmplx, iterations uint8) color.Color {
	const contrast = 5

	v := cmplxbigrat.New(complex(0, 0))
	for n := uint8(0); n < iterations; n++ {
		v = cmplxbigrat.Add(cmplxbigrat.Mul(v, v), z)
		// SqAbs returns abs*abs
		if cmplxbigrat.SqAbs(v).Cmp(new(big.Rat).SetFloat64(4)) == 1 {
			return color.YCbCr{contrast * n, 255 - contrast*n, contrast * n}
		}
	}
	return color.Black
}
