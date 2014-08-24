package main

import (
	"log"
	"runtime"

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
	runtime.LockOSThread()

	var err error
	if !glfw.Init() {
		log.Fatalf("Failed Init\n")
		return
	}

	defer glfw.Terminate()
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 0)
	//glfw.WindowHint(glfw.OpenglForwardCompatible, glfw.True)
	//glfw.WindowHint(glfw.OpenglProfile, glfw.OpenglCoreProfile)
	window, err := glfw.CreateWindow(Width, Height, "Window Title", nil, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
		return
	}
	defer window.Destroy()
	window.SetSizeCallback(onResize)
	window.SetKeyCallback(onKey)

	window.MakeContextCurrent()
	glfw.SwapInterval(1)

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
