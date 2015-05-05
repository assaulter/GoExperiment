// エクササイズ60
// Imageインタフェースを実装し、ShowImageをコール出来るようにする。
package main

import (
	"image"
	"image/color"

	"code.google.com/p/go-tour/pic"
)

/*
type Image interface {
        // ColorModel returns the Image's color model.
        ColorModel() color.Model
        // Bounds returns the domain for which At can return non-zero color.
        // The bounds do not necessarily contain the point (0, 0).
        Bounds() Rectangle
        // At returns the color of the pixel at (x, y).
        // At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
        // At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
        At(x, y int) color.Color
}*/

type Image struct{}

func (i *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i *Image) Bounds() image.Rectangle {
	return image.Rectangle{
		image.Point{0, 0},
		image.Point{200, 200},
	}
}

func (i *Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x % 256), uint8(y % 256), uint8((x * y) % 256), 255}
}

func main() {
	m := Image{}
	pic.ShowImage(&m)
}
