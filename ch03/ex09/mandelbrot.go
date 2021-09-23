package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"net/url"
	"strconv"
)

const (
	width, height = 1024, 1024
)

type imgParam struct {
	x    float64
	y    float64
	zoom int
}

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		param, err := parseQuery(r.URL)
		if err != nil {
			log.Print(err)
		}
		draw(w, param)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func parseQuery(url *url.URL) (imgParam, error) {
	defaultX := -2.0
	defaultY := -2.0
	defaultZoom := 1

	q := url.Query()
	if q == nil {
		return imgParam{
			x:    defaultX,
			y:    defaultY,
			zoom: defaultZoom,
		}, nil
	}

	x, err := strconv.ParseFloat(q.Get("x"), 64)
	if err != nil {
		x = defaultX
	}

	y, err := strconv.ParseFloat(q.Get("y"), 64)
	if err != nil {
		y = defaultY
	}

	zoom, err := strconv.Atoi(q.Get("zoom"))
	if err != nil {
		zoom = defaultZoom
	}

	return imgParam{
		x:    x,
		y:    y,
		zoom: zoom,
	}, nil
}

func draw(out io.Writer, param imgParam) {
	xmin := param.x
	ymin := param.y
	xmax := float64(xmin) + 4/float64(param.zoom)
	ymax := float64(ymin) + 4/float64(param.zoom)
	fmt.Printf("%f to %f, %f to %f\n", xmin, xmax, ymin, ymax)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrotCmplx128(z, 255))
		}
	}
	png.Encode(out, img)
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
