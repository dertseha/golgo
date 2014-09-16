package util

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"path"
)

/**
 * GetMatch returns the percentage (0.0 .. 1.0) of how much the passed image
 * equals to the one referenced via file name.
 */
func GetMatch(refName string, input image.Image) float32 {
	refFile, _ := os.Open(refName)
	refImg, _ := png.Decode(refFile)
	refFile.Close()

	dx := refImg.Bounds().Dx()
	dy := refImg.Bounds().Dy()
	equal := 0
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			if DoesPixelMatch(refImg.At(x, y), x, y, input, 5*0x100) {
				equal++
			}
		}
	}

	return float32(equal) / float32(dx*dy)
}

func DoesPixelMatch(pixel color.Color, x, y int, ref image.Image, delta int) bool {
	result := DoesColorMatch(pixel, ref.At(x, y), delta)

	if !result {
		rows := make([]int, 0, 2)
		columns := make([]int, 0, 2)
		if y > 0 {
			rows = append(rows, y-1)
		}
		if (y + 1) < ref.Bounds().Dy() {
			rows = append(rows, y+1)
		}
		if x > 0 {
			columns = append(columns, x-1)
		}
		if (x + 1) < ref.Bounds().Dx() {
			columns = append(columns, x+1)
		}

		for rowIndex := 0; !result && rowIndex < len(rows); rowIndex++ {
			for columnIndex := 0; !result && columnIndex < len(columns); columnIndex++ {
				result = DoesColorMatch(pixel, ref.At(columns[columnIndex], rows[rowIndex]), delta)
			}
		}
	}

	return result
}

func DoesColorMatch(p1 color.Color, p2 color.Color, delta int) bool {
	r1, g1, b1, a1 := p1.RGBA()
	r2, g2, b2, a2 := p2.RGBA()

	return math.Abs(float64(r1)-float64(r2)) <= float64(delta) &&
		math.Abs(float64(g1)-float64(g2)) <= float64(delta) &&
		math.Abs(float64(b1)-float64(b2)) <= float64(delta) &&
		math.Abs(float64(a1)-float64(a2)) <= float64(delta)
}

func SaveImage(name string, img image.Image) {
	os.MkdirAll(path.Dir(name), os.FileMode(0755))
	writer, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer writer.Close()

	png.Encode(writer, img)
}
