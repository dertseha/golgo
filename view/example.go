package view

import (
	gles "github.com/dertseha/golgo/gles2"
)

func DrawExampleScene(gl gles.OpenGl) {
	gl.ClearColor(0.3, 0.5, 0.1, 1)
	gl.Clear(gles.COLOR_BUFFER_BIT | gles.DEPTH_BUFFER_BIT)
}
