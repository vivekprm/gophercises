package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

// https://github.com/golang/go/tree/master/src/image
// User sourcegraph to serach implmentation of Image interface
func main() {
	data := []int{10, 33, 73, 64}
	w, h := len(data)*60+10, 100
	r := image.Rect(0, 0, w, h)
	img := image.NewRGBA(r)
	bg := image.NewUniform(color.RGBA{240, 240, 240, 255})
	draw.Draw(img, r, bg, image.Point{0, 0}, draw.Src)

	// ImageDrawingDraw(img, r, data)
	ImageDrawingMaskedDraw(img, r, data, w, h)
	// ImageDrawingPixel(img, data, w, h)
}
func DrawMask(w, h int) *image.RGBA {
	mask := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			mask.Set(x, y, color.RGBA{
				R: uint8((x + y) & 255),
				G: uint8((x + y) << 1 & 255),
				B: uint8((x + y) << 2 & 255),
				A: uint8((255 * (100 - y)) / 100),
			})
		}
	}
	return mask
}
func ImageDrawingMaskedDraw(img *image.RGBA, r image.Rectangle, data []int, w, h int) {
	mask := DrawMask(w, h)
	for i, dp := range data {
		x0, y0 := i*60+10, 100-dp
		x1, y1 := (i+1)*60-1, 100
		bar := image.Rect(x0, y0, x1, y1)
		grey := image.NewUniform(color.RGBA{180, 180, 180, 255})
		draw.Draw(img, bar, grey, image.Point{0, 0}, draw.Src)
		red := image.NewUniform(color.RGBA{250, 0, 0, 255})
		draw.DrawMask(img, bar, red, image.Point{0, 0}, mask, image.Point{x0, y0}, draw.Over)
	}
	savePngImage(img, "image3.png")
}

func ImageDrawingDraw(img *image.RGBA, r image.Rectangle, data []int) {
	grey := image.NewUniform(color.RGBA{240, 240, 240, 255})
	red := image.NewUniform(color.RGBA{250, 180, 180, 255})

	draw.Draw(img, r, grey, image.Point{0, 0}, draw.Src)
	for i, dp := range data {
		x0, y0 := i*60+10, 100-dp
		x1, y1 := (i+1)*60-1, 100
		bar := image.Rect(x0, y0, x1, y1)
		draw.Draw(img, bar, red, image.Point{0, 0}, draw.Src)
	}
	savePngImage(img, "image2.png")
}

func ImageDrawingPixel(img *image.RGBA, data []int, w, h int) {
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			img.Set(x, y, color.RGBA{255, 255, 255, 255})
		}
	}

	for i, dp := range data {
		for x := i*60 + 10; x < i*60+60; x++ {
			for y := 100; y >= (100 - dp); y-- {
				img.Set(x, y, color.RGBA{180, 180, 250, 255})
			}
		}
	}
	savePngImage(img, "image1.png")
}

func savePngImage(img *image.RGBA, file string) {
	f, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := png.Encode(f, img); err != nil {
		log.Fatal(err)
	}
}
