package wrapper

import (
	"github.com/gopherjs/webgl"
)

type WebGl struct {
	gl *webgl.Context
}

func (gl *WebGl) Clear(mask uint) {
	gl.gl.Clear(int(mask))
}

func (gl *WebGl) ClearColor(red float32, green float32, blue float32, alpha float32) {
	gl.gl.ClearColor(red, green, blue, alpha)
}
