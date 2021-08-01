package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	rect image.Rectangle
}

func (i Image) Bounds() image.Rectangle {
	return i.rect
}

func (Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (Image) At(x, y int) color.Color {
	v := uint8(x ^ 1200 + y ^ 450)
	return color.RGBA{v, v, 120, 120}
}

func main() {
	m := Image{image.Rect(0, 0, 256, 128)}
	pic.ShowImage(m)
}
