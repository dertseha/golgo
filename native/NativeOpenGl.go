package native

import (
	"github.com/go-gl/gl"
)

type NativeOpenGl struct{}

func (*NativeOpenGl) Clear(mask uint) {
	gl.Clear(gl.GLbitfield(mask))
}

func (*NativeOpenGl) ClearColor(red float32, green float32, blue float32, alpha float32) {
	gl.ClearColor(gl.GLclampf(red), gl.GLclampf(green), gl.GLclampf(blue), gl.GLclampf(alpha))
}
