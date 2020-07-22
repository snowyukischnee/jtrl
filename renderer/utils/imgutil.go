package utils

import (
	. "image"
	"image/color"
	"image/png"
	"os"
	"renderer/core"
)

type ImageWrapper struct {
	Img *RGBA
	Width int
	Height int
}

func NewImage(width int, height int) *ImageWrapper {
	return &ImageWrapper{
		Img:    NewRGBA(Rectangle{Min: Point{X: 0, Y: 0}, Max: Point{X: width, Y: height}}),
		Width:  width,
		Height: height,
	}
}

func (img *ImageWrapper) WriteImage(path string) {
	f, err := os.Create(path)
	CheckError(err)
	err = png.Encode(f, img.Img)
	CheckError(err)
}

func (img *ImageWrapper) Paint(x int, y int, v *core.Point3)  {
	img.Img.Set(x, y, color.RGBA{
		R: uint8(v.X),
		G: uint8(v.Y),
		B: uint8(v.Z),
		A: 0xff,
	})
}

func (img *ImageWrapper) PaintC(x int, y int, c color.Color)  {
	img.Img.Set(x,y, c)
}