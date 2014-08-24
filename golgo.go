package main

import (
	"log"

	gogl "github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"

	"github.com/dertseha/golgo/native"
	"github.com/dertseha/golgo/view"
)

const (
	Width  = 640
	Height = 480
)

var running bool

func main() {
	var err error
	if !glfw.Init() {
		log.Fatalf("Failed Init\n")
		return
	}

	defer glfw.Terminate()

	window, err := glfw.CreateWindow(Width, Height, "Window Title", nil, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
		return
	}

	window.MakeContextCurrent()
	window.SetSizeCallback(onResize)
	window.SetKeyCallback(onKey)

	gogl.Init()
	wrappedGl := native.CreateGles2Wrapper()

	example := view.NewNeheExample02(wrappedGl, Width, Height)
	example.Init()

	running = true
	for running && !window.ShouldClose() {
		example.DrawScene()

		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func onResize(window *glfw.Window, w int, h int) {

}

func onKey(window *glfw.Window, key glfw.Key, unknown int, action glfw.Action, modifier glfw.ModifierKey) {
	switch key {
	case glfw.KeyEscape:
		running = false
	}
}
