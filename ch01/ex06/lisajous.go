package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{
	color.Black,
	color.RGBA{255, 0, 0, 1},
	color.RGBA{255, 165, 0, 1},
	color.RGBA{255, 255, 0, 1},
	color.RGBA{0, 128, 0, 1},
	color.RGBA{0, 0, 255, 1},
	color.RGBA{75, 0, 130, 1},
	color.RGBA{238, 130, 238, 1},
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lisajous(os.Stdout)
}

func createPalette(nframes int) []color.Color {
	var palette []color.Color
	for i := 0; i < nframes; i++ {
		palette = append(palette, color.RGBA{0, 255, 0, uint8(255 / nframes * i)})
	}
	return palette
}

func fillRect(img *image.Paletted, col color.Color) {
	rect := img.Rect
	for h := rect.Min.Y; h < rect.Max.Y; h++ {
		for v := rect.Min.X; v < rect.Max.X; v++ {
			img.Set(v, h, col)
		}
	}
}

func lisajous(out io.Writer) {
	const (
		cycles  = 5
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
		fillRect(img, color.Black)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			colorIdx := uint8(int(math.Abs(x)*7) + 1)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIdx)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
