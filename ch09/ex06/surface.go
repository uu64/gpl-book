package main

import (
	"fmt"
	"image/color"
	"io"
	"math"
	"os"
	"sync"
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

type point struct {
	x, y float64
}

type polygon struct {
	clr        *color.NRGBA
	a, b, c, d *point
}

func main() {
	concurrentDraw(os.Stdout, 2)
}

func concurrentDraw(w io.Writer, concurrency int) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	var wg sync.WaitGroup

	slots := make(chan struct{}, concurrency)
	ch := make(chan *polygon)

	calc := func(i, j int, channel chan<- *polygon, slots <-chan struct{}) {
		ax, ay, ah, aerr := corner(i+1, j)
		bx, by, bh, berr := corner(i, j)
		cx, cy, ch, cerr := corner(i, j+1)
		dx, dy, dh, derr := corner(i+1, j+1)
		if aerr != nil || berr != nil || cerr != nil || derr != nil {
			channel <- nil
		} else {
			clr := calcColor((ah + bh + ch + dh) / 4)

			channel <- &polygon{
				clr: &clr,
				a:   &point{ax, ay},
				b:   &point{bx, by},
				c:   &point{cx, cy},
				d:   &point{dx, dy},
			}
		}
		<-slots
	}

	draw := func(channel <-chan *polygon) {
		defer wg.Done()

		for i := 0; i < cells; i++ {
			for j := 0; j < cells; j++ {
				if p := <-channel; p != nil {
					fmt.Fprintf(w, "<polygon style='fill:#%02x%02x%02x%02x' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
						p.clr.R, p.clr.G, p.clr.B, p.clr.A, p.a.x, p.a.y, p.b.x, p.b.y, p.c.x, p.c.y, p.d.x, p.d.y)
				}
			}
		}
	}

	wg.Add(1)
	go draw(ch)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			slots <- struct{}{}
			go calc(i, j, ch, slots)
		}
	}

	wg.Wait()
	fmt.Fprintln(w, "</svg>")
}

//lint:ignore U1000 this function is used in the test
func draw(w io.Writer) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ah, aerr := corner(i+1, j)
			bx, by, bh, berr := corner(i, j)
			cx, cy, ch, cerr := corner(i, j+1)
			dx, dy, dh, derr := corner(i+1, j+1)
			if aerr != nil || berr != nil || cerr != nil || derr != nil {
				continue
			}

			clr := calcColor((ah + bh + ch + dh) / 4)

			fmt.Fprintf(w, "<polygon style='fill:#%02x%02x%02x%02x' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				clr.R, clr.G, clr.B, clr.A, ax, ay, bx, by, cx, cy, dx, dy)
		}

	}
	fmt.Fprintln(w, "</svg>")
}

func calcColor(height float64) color.NRGBA {
	hmin := -0.3
	hmax := 1.0
	level := (height - hmin) / ((hmax - hmin) / 255)
	return color.NRGBA{
		R: uint8(level),
		G: uint8(0),
		B: uint8(255 - level),
		A: uint8(255),
	}
}

func corner(i, j int) (float64, float64, float64, error) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	if math.IsNaN(z) || math.IsInf(z, 0) {
		return 0, 0, 0, fmt.Errorf("the value is invalid: %g", z)
	}

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z, nil
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
