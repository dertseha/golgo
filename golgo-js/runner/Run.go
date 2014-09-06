package runner

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/webgl"

	"github.com/dertseha/golgo/app"
	"github.com/dertseha/golgo/golgo-js/wrapper"
)

func Run(application app.Application, param app.ApplicationParameter) {
	document := js.Global.Get("document")
	document.Set("title", param.Title())

	canvas := document.Call("createElement", "canvas")
	canvas.Set("width", param.Width())
	canvas.Set("height", param.Height())
	document.Get("body").Call("appendChild", canvas)

	attrs := webgl.DefaultAttributes()
	attrs.Alpha = false

	glContext, err := webgl.NewContext(canvas, attrs)
	if err != nil {
		panic(err)
	}

	gl := wrapper.CreateGles2Wrapper(glContext)

	application.Init(gl, param.Width(), param.Height())

	requestAnimation(application)
}

func requestAnimation(application app.Application) {
	window := js.Global.Get("window")
	type indirecterType struct {
		render func()
	}
	var indirecter indirecterType

	indirecter.render = func() {
		window.Call("requestAnimationFrame", indirecter.render)
		application.Render()
	}
	indirecter.render()
}
