package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

const (
	cells   = 100
	xyrange = 30.0
	angle   = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

type svgParams struct {
	width  int
	height int
	color  string
}

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		params, err := parseQuery(r.URL)
		if err != nil {
			log.Print(err)
		}

		draw(w, params)
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func parseQuery(url *url.URL) (svgParams, error) {
	defaultWidth := 600
	defaultHeight := 320
	defaultColor := "white"
	q := url.Query()
	if q == nil {
		return svgParams{
			width:  defaultWidth,
			height: defaultHeight,
			color:  defaultColor,
		}, nil
	}

	width, err := strconv.Atoi(q.Get("width"))
	if err != nil {
		width = defaultWidth
	}

	height, err := strconv.Atoi(q.Get("height"))
	if err != nil {
		height = defaultHeight
	}

	color := q.Get("color")
	if len(color) == 0 {
		color = defaultColor
	}

	return svgParams{
		width:  width,
		height: height,
		color:  color,
	}, nil
}

func draw(out io.Writer, params svgParams) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", params.width, params.height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, aerr := corner(i+1, j, params)
			bx, by, berr := corner(i, j, params)
			cx, cy, cerr := corner(i, j+1, params)
			dx, dy, derr := corner(i+1, j+1, params)
			if aerr != nil || berr != nil || cerr != nil || derr != nil {
				continue
			}
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:%s' />\n",
				ax, ay, bx, by, cx, cy, dx, dy, params.color)
		}

	}
	fmt.Fprintf(out, "</svg>")
}

func corner(i, j int, params svgParams) (float64, float64, error) {
	xyscale := float64(params.width) / 2 / xyrange
	zscale := float64(params.height) * 0.4
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)
	if math.IsNaN(z) || math.IsInf(z, 0) {
		return 0, 0, fmt.Errorf("the value is invalid: %g", z)
	}

	sx := float64(params.width/2) + (x-y)*cos30*xyscale
	sy := float64(params.height/2) + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, nil
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
