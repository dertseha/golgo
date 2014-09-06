package test

import (
	gles "github.com/dertseha/golgo/gles2"
)

type clearBackgroundApplication struct {
	gl gles.OpenGl
}

func NewClearBackgroundApplication() *clearBackgroundApplication {
	return &clearBackgroundApplication{}
}

func (app *clearBackgroundApplication) Init(gl gles.OpenGl, width, height int) {
	app.gl = gl

	app.gl.ClearColor(0.3, 0.5, 0.1, 1.0)
}

func (app *clearBackgroundApplication) Resize(width, height int) {

}

func (app *clearBackgroundApplication) Render() {
	app.gl.Clear(gles.COLOR_BUFFER_BIT | gles.DEPTH_BUFFER_BIT)
}
