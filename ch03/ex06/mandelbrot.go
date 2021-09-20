package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < 2*height; py++ {
		y1 := float64(py)/height*(ymax-ymin) + ymin
		y2 := float64(py+1)/height*(ymax-ymin) + ymin
		for px := 0; px < 2*width; px++ {
			x1 := float64(px)/width*(xmax-xmin) + xmin
			x2 := float64(px+1)/width*(xmax-xmin) + xmin

			// supersampling
			c1 := mandelbrot(complex(x1, y1))
			c2 := mandelbrot(complex(x1, y2))
			c3 := mandelbrot(complex(x2, y1))
			c4 := mandelbrot(complex(x2, y2))
			color := color.YCbCr{
				uint8((int(c1.Y) + int(c2.Y) + int(c3.Y) + int(c4.Y)) / 4),
				uint8((int(c1.Cb) + int(c2.Cb) + int(c3.Cb) + int(c4.Cb)) / 4),
				uint8((int(c1.Cr) + int(c2.Cr) + int(c3.Cr) + int(c4.Cr)) / 4),
			}
			img.Set(px, py, color)
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.YCbCr {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.YCbCr{contrast * n, 255 - contrast*n, contrast * n}
		}
	}
	y, cb, cr := color.RGBToYCbCr(0, 0, 0)
	return color.YCbCr{y, cb, cr}
}
