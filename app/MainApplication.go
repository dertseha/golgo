package app

import (
	gles "github.com/dertseha/golgo/gles2"
)

type mainApplication struct {
	gl gles.OpenGl
}

func NewMainApplication() Application {
	return &mainApplication{}
}

func (app *mainApplication) Init(gl gles.OpenGl, width, height int) {
	app.gl = gl

	app.gl.ClearColor(0.0, 0.0, 0.0, 1.0)
}

func (app *mainApplication) Resize(width, height int) {

}

func (app *mainApplication) Render() {
	app.gl.Clear(gles.COLOR_BUFFER_BIT | gles.DEPTH_BUFFER_BIT)
}
