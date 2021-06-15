package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}
func (img Image) Bounds() image.Rectangle {
	// First method.
	// return image.Rectangle{
	// 	Min: image.Point{X: 0, Y: 0},
	// 	Max: image.Point{X: 100, Y: 100},
	// }
	// - - -

	// Second method.
	return image.Rect(0, 0, 256, 256)

}
func (img Image) At(x, y int) color.Color {
	// r, g := uint8(0), uint8(0)
	// return color.RGBA{R: r, G: g, B: 255, A: 255}

	v := uint8(x ^ y) // x ^ y 按位异或
	return color.RGBA{R: v, G: v, B: 255, A: 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
