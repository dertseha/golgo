package core

import (
	gles "github.com/dertseha/golgo/gles2"
)

type MainApplication struct {
	gl gles.OpenGl
}

func NewMainApplication() *MainApplication {
	return &MainApplication{}
}

func (app *MainApplication) Init(gl gles.OpenGl, width, height int) {
	app.gl = gl

	app.gl.ClearColor(0.0, 0.0, 0.0, 1.0)
}

func (app *MainApplication) Resize(width, height int) {

}

func (app *MainApplication) Render() {
	app.gl.Clear(gles.COLOR_BUFFER_BIT | gles.DEPTH_BUFFER_BIT)
}
