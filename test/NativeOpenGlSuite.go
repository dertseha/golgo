package test

import (
	gogl "github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"

	gles "github.com/dertseha/golgo/gles2"
	"github.com/dertseha/golgo/native"
	"github.com/dertseha/golgo/util"

	testUtil "github.com/dertseha/golgo/test/util"

	check "gopkg.in/check.v1"
)

type NativeOpenGlSuite struct {
	width  int
	height int

	window *glfw.Window
	gl     gles.OpenGl
}

func (suite *NativeOpenGlSuite) SetUpTest(c *check.C) {
	glfw.Init()
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 0)

	window, _ := glfw.CreateWindow(suite.width, suite.height, "Native OpenGL Test", nil, nil)
	window.MakeContextCurrent()
	glfw.SwapInterval(1)

	gogl.Init()

	suite.window = window
	suite.gl = native.CreateGles2Wrapper()
}

func (suite *NativeOpenGlSuite) TearDownTest(c *check.C) {
	suite.window.Destroy()
	glfw.Terminate()
}

func (suite *NativeOpenGlSuite) ThenScreenShouldMatchReference(c *check.C, refName string) {
	liveImg, glError := util.ReadPixels(suite.gl, 0, 0, suite.width, suite.height)
	c.Assert(glError, check.Equals, gles.NO_ERROR)

	testUtil.SaveImage("screenshots/native/"+refName, liveImg)
	confidence := testUtil.GetMatch("resources/"+refName, liveImg)

	c.Check(confidence >= 0.99, check.Equals, true, check.Commentf("Confidence == %f", confidence))
}
