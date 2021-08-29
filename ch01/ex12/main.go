package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var palette = []color.Color{color.White, color.Black, color.RGBA{0, 255, 0, 1}}

const (
	whiteIndex = 0
	blackIndex = 1
	greenIndex = 2
)

func fillRect(img *image.Paletted, col color.Color) {
	rect := img.Rect
	for h := rect.Min.Y; h < rect.Max.Y; h++ {
		for v := rect.Min.X; v < rect.Max.X; v++ {
			img.Set(v, h, col)
		}
	}
}

func lissajous(out io.Writer, cycles int) {
	log.Printf("cycles: %d\n", cycles)

	const (
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		fillRect(img, palette[blackIndex])
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func parseQuery(url *url.URL) (int, error) {
	q := url.Query()
	if q == nil {
		return 0, nil
	}

	v := q.Get("cycles")
	if len(v) == 0 {
		return 0, nil
	}

	cycles, err := strconv.Atoi(v)
	if err != nil {
		return 0, err
	}
	return cycles, nil
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	handler := func(w http.ResponseWriter, r *http.Request) {
		cycles, err := parseQuery(r.URL)
		if err != nil {
			log.Print(err)
		}

		lissajous(w, cycles)
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
