package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

type MyColor struct{}

func (c MyColor) RGBA() (r, g, b, a uint32) {
	return 10, 10, 10, 0
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 100)
}

func (img Image) At(x, y int) color.Color {
	return MyColor{}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
