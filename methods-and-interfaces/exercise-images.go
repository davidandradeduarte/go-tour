package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	Height int
	Width  int
}

func main() {
	m := Image{50, 50}
	pic.ShowImage(m)
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.Width, i.Height)
}

func (i Image) At(x, y int) color.Color {
	c := uint8(x ^ y)
	return color.RGBA{c, c, 25, 25}
}
