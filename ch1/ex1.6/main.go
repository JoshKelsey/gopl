package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.RGBA{0, 0, 0, 1}, color.RGBA{148, 0, 211, 1}, color.RGBA{75, 0, 130, 1},
	color.RGBA{75, 0, 130, 1}, color.RGBA{0, 0, 255, 1},
	color.RGBA{0, 255, 0, 1}, color.RGBA{255, 255, 0, 1},
	color.RGBA{255, 127, 0, 1}, color.RGBA{255, 0, 0, 1},
}

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
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
		var index uint8
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), index)
			if index != 7 {
				index++
			} else {
				index -= 7
			}
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
