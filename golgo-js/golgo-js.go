package main

// Some resources:
//
// http://learningwebgl.com/lessons/
// http://opengles-book.com/index.html
// https://github.com/danginsburg/opengles3-book
//
// Even more:
// http://blog.cmgresearch.com/2010/10/05/NeHe-OpenGL-tutorial-2%2C-3-and-4-ported-to-ES-2.0/
// view-source:http://learningwebgl.com/lessons/lesson02/index.html
//
// https://github.com/ungerik/go3d
// http://godoc.org/github.com/gopherjs/gopherjs/js
//
// http://www.html5rocks.com/en/tutorials/webgl/webgl_fundamentals/
// https://www.khronos.org/registry/webgl/specs/1.0/#5.13
// http://www.khronos.org/registry/gles/api/GLES2/gl2.h
//

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/webgl"

	"github.com/dertseha/golgo/golgo-js/wrapper"
	"github.com/dertseha/golgo/view"
)

const (
	Width  = 640
	Height = 480
)

func main() {
	document := js.Global.Get("document")
	canvas := document.Call("createElement", "canvas")
	canvas.Set("width", Width)
	canvas.Set("height", Height)
	document.Get("body").Call("appendChild", canvas)

	attrs := webgl.DefaultAttributes()
	attrs.Alpha = false

	gl, err := webgl.NewContext(canvas, attrs)
	if err != nil {
		js.Global.Call("alert", "Error: "+err.Error())
	}

	wrappedGl := wrapper.CreateGles2Wrapper(gl)

	example := view.NewNeheExample02(wrappedGl, Width, Height)
	example.Init()
	example.DrawScene()
}
