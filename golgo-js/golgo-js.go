package main

// Some resources:
//
// http://learningwebgl.com/lessons/
// http://opengles-book.com/index.html
// https://github.com/danginsburg/opengles3-book

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

	view.DrawExampleScene(&wrapper.WebGl{gl})
}
