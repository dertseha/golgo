package native

import (
	"path"

	gogl "github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"

	gles "github.com/dertseha/golgo/gles2"
	"github.com/dertseha/golgo/runner/native"
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

func NewNativeOpenGlSuite(width, height int) NativeOpenGlSuite {
	return NativeOpenGlSuite{width: width, height: height}
}

func (suite *NativeOpenGlSuite) Width() int {
	return suite.width
}

func (suite *NativeOpenGlSuite) Height() int {
	return suite.height
}

func (suite *NativeOpenGlSuite) OpenGl() gles.OpenGl {
	return suite.gl
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
	screenshotPath := path.Join(testUtil.GetProjectRootPath(), "test", "screenshots", "native", c.TestName()+"."+refName)
	referencePath := path.Join(testUtil.GetProjectRootPath(), "test", "resources", refName)
	liveImg, glError := util.ReadPixels(suite.gl, 0, 0, suite.width, suite.height)
	c.Assert(glError, check.Equals, gles.NO_ERROR)

	testUtil.SaveImage(screenshotPath, liveImg)
	confidence := testUtil.GetMatch(referencePath, liveImg)

	c.Check(confidence >= 0.99, check.Equals, true, check.Commentf("Confidence == %f", confidence))
}
