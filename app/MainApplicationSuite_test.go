package app

import (
	"github.com/dertseha/golgo/test/native"

	check "gopkg.in/check.v1"
)

type MainApplicationSuite struct {
	native.NativeOpenGlSuite
}

var _ = check.Suite(&MainApplicationSuite{native.NewNativeOpenGlSuite(640, 480)})

func (suite *MainApplicationSuite) TestInitialRender(c *check.C) {
	app := NewMainApplication()

	app.Init(suite.OpenGl(), suite.Width(), suite.Height())

	app.Render()

	suite.ThenScreenShouldMatchReference(c, "MainApp_InitialRender.png")
}
