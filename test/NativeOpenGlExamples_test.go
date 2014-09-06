package test

import (
	check "gopkg.in/check.v1"
)

type NativeOpenGlExamplesSuite struct {
	NativeOpenGlSuite
}

var _ = check.Suite(&NativeOpenGlExamplesSuite{NativeOpenGlSuite{width: 640, height: 480}})

func (suite *NativeOpenGlExamplesSuite) TestClearBackground(c *check.C) {
	app := &clearBackgroundApplication{}

	app.Init(suite.gl, suite.width, suite.height)

	app.Render()

	suite.ThenScreenShouldMatchReference(c, "ClearBackground.png")
}

func (suite *NativeOpenGlExamplesSuite) TestNeheExample02(c *check.C) {
	app := NewNeheExample02Application()

	app.Init(suite.gl, suite.width, suite.height)

	app.Render()

	suite.ThenScreenShouldMatchReference(c, "NeheExample02.png")
}
