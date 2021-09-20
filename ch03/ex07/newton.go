package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -1, -1, +1, +1
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func newton(z0 complex128) color.YCbCr {
	const iterations = 255
	const contrast = 15
	const tolerance = 1.e-8

	// newton method
	// f(z) = z^4 - 1
	// f'(z) = 4z^3
	roots := []complex128{
		complex(1, 0),
		complex(-1, 0),
		complex(0, 1),
		complex(0, -1),
	}

	z := z0
	for n := uint8(0); n < iterations; n++ {
		dz := (z*z*z*z - 1) / (4 * z * z * z)
		idx, distance := getNearestRoot(z, roots)
		if distance < tolerance {
			switch idx {
			case 0:
				return color.YCbCr{contrast * n, 0, 0}
			case 1:
				return color.YCbCr{contrast * n, 0, 255}
			case 2:
				return color.YCbCr{contrast * n, 255, 0}
			case 3:
				return color.YCbCr{contrast * n, 255, 255}
			}
		}
		z -= dz
	}
	y, cb, cr := color.RGBToYCbCr(0, 0, 0)
	return color.YCbCr{y, cb, cr}
}

func getNearestRoot(c complex128, roots []complex128) (int, float64) {
	min := math.MaxFloat64
	minIdx := 0
	for i, r := range roots {
		d := cmplx.Abs(r - c)
		if d < min {
			min = d
			minIdx = i
		}
	}
	return minIdx, min
}
