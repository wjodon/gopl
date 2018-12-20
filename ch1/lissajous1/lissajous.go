// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"golang.org/x/image/colornames"
)

var palette = []color.Color{
	color.Black,
	colornames.Lightgreen,
	colornames.Lightblue,
	colornames.Lightcoral,
	colornames.Lightcyan,
	colornames.Lightgoldenrodyellow,
	colornames.Lightpink,
	colornames.Lightsalmon,
	colornames.Lightseagreen,
	colornames.Lightskyblue,
	colornames.Lightslategray,
	colornames.Lightsteelblue,
	colornames.Lightyellow }


func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles 	= 5		// number of complete x oscillator revolutions
		res 	= 0.001	// angular resolution
		size	= 100	// image canvas covers [-size..+size]
		nframes	= 64	// number of animation frames
		delay	= 8		// delay between frames in 10ms units
		)

	freq := rand.Float64() * 3.0 // relative frequency of y oxcillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	foregroundIndex := uint8(1)
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2 * size + 1, 2 * size + 1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles * 2 * math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t * freq + phase)
			img.SetColorIndex(size + int(x * size + 0.5), size + int(y * size + 0.5),
				foregroundIndex)
		}
		if int(phase * 10) % 5 == 0 {
			foregroundIndex = uint8(rand.Intn(12) + 1)
		} 
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}	