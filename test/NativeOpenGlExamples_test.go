package test

import (
	"github.com/dertseha/golgo/view"

	check "gopkg.in/check.v1"
)

type NativeOpenGlExamplesSuite struct {
	NativeOpenGlSuite
}

var _ = check.Suite(&NativeOpenGlExamplesSuite{NativeOpenGlSuite{width: 640, height: 480}})

func (suite *NativeOpenGlExamplesSuite) TestNeheExample02(c *check.C) {
	example := view.NewNeheExample02(suite.gl, suite.width, suite.height)
	example.Init()

	example.DrawScene()

	suite.ThenScreenShouldMatchReference(c, "NeheExample02.png")
}
