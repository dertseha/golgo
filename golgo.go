package main

import (
	"log"
	"math"

	"github.com/go-gl/gl"
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

	initGL()
	onResize(window, Width, Height)

	running = true
	for running && !window.ShouldClose() {

		view.DrawExampleScene(&native.NativeOpenGl{})

		window.SwapBuffers()
		glfw.PollEvents()
	}

}

func onResize(window *glfw.Window, w int, h int) {
	if h == 0 {
		h = 1
	}

	gl.Viewport(0, 0, w, h)
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	setPerspective(90.0/2.0, float64(w)/float64(h), 0.1, 100.0)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
}

/**
 * Code to avoid importing GLU. Problem is though that also gl.Frustum was declared
 * deprecated in GL 3.1
 * See http://stackoverflow.com/questions/2417697/gluperspective-was-removed-in-opengl-3-1-any-replacements
 */
func setPerspective(fieldOfView float64, aspect float64, zNear float64, zFar float64) {
	fH := math.Tan(fieldOfView/360.0*3.14159) * zNear
	fW := fH * aspect

	gl.Frustum(-fW, fW, -fH, fH, zNear, zFar)
}

func onKey(window *glfw.Window, key glfw.Key, unknown int, action glfw.Action, modifier glfw.ModifierKey) {
	switch key {
	case glfw.KeyEscape:
		running = false
	}
}

func initGL() {
	gl.ShadeModel(gl.SMOOTH)
	//gl.ClearColor(0, 0, 0, 0)
	gl.ClearDepth(1)
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LEQUAL)
	gl.Hint(gl.PERSPECTIVE_CORRECTION_HINT, gl.NICEST)
}
