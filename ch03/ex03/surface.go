package main

import (
	"fmt"
	"image/color"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, aclr, aerr := corner(i+1, j)
			bx, by, bclr, berr := corner(i, j)
			cx, cy, cclr, cerr := corner(i, j+1)
			dx, dy, dclr, derr := corner(i+1, j+1)
			if aerr != nil || berr != nil || cerr != nil || derr != nil {
				continue
			}
			clr := color.NRGBA{
				R: uint8((int(aclr.R) + int(bclr.R) + int(cclr.R) + int(dclr.R)) / 4),
				G: uint8((int(aclr.G) + int(bclr.G) + int(cclr.G) + int(dclr.G)) / 4),
				B: uint8((int(aclr.B) + int(bclr.B) + int(cclr.B) + int(dclr.B)) / 4),
				A: uint8((int(aclr.A) + int(bclr.A) + int(cclr.A) + int(dclr.A)) / 4),
			}

			fmt.Printf("<polygon style='fill:#%02x%02x%02x%02x' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				clr.R, clr.G, clr.B, clr.A, ax, ay, bx, by, cx, cy, dx, dy)
		}

	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, color.NRGBA, error) {
	color := color.NRGBA{}
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	if math.IsNaN(z) || math.IsInf(z, 0) {
		return 0, 0, color, fmt.Errorf("the value is invalid: %g", z)
	}
	zmin := -0.3
	zmax := 1.0
	color.R = uint8((z - zmin) / ((zmax - zmin) / 255))
	color.G = 0
	color.B = 255 - uint8((z-zmin)/((zmax-zmin)/255))
	color.A = 255

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, color, nil
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
