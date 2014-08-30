package util

import (
	"image"
	"image/color"

	gles "github.com/dertseha/golgo/gles2"
)

const (
	bytesPerPixel = 4
)

type rawImage struct {
	data   []byte
	width  int
	height int
}

func (raw *rawImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (raw *rawImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, raw.width, raw.height)
}

func (raw *rawImage) At(x, y int) color.Color {
	offset := (x + ((raw.height - 1 - y) * raw.width)) * bytesPerPixel
	result := color.RGBA{
		R: raw.data[offset],
		G: raw.data[offset+1],
		B: raw.data[offset+2],
		A: raw.data[offset+3]}

	return result
}

func ReadPixels(gl gles.OpenGl, x, y, width, height int) (img image.Image, glError uint) {
	data := make([]byte, width*height*bytesPerPixel)
	img = &rawImage{
		data:   data,
		width:  width,
		height: height}

	gl.ReadPixels(x, y, width, height, gles.RGBA, gles.UNSIGNED_BYTE, data)
	glError = gl.GetError()

	return
}
