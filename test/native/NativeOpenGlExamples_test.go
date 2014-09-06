package native

import (
	"github.com/dertseha/golgo/test"

	check "gopkg.in/check.v1"
)

type NativeOpenGlExamplesSuite struct {
	NativeOpenGlSuite
}

var _ = check.Suite(&NativeOpenGlExamplesSuite{NewNativeOpenGlSuite(640, 480)})

func (suite *NativeOpenGlExamplesSuite) TestClearBackground(c *check.C) {
	app := test.NewClearBackgroundApplication()

	app.Init(suite.OpenGl(), suite.Width(), suite.Height())

	app.Render()

	suite.ThenScreenShouldMatchReference(c, "ClearBackground.png")
}

func (suite *NativeOpenGlExamplesSuite) TestNeheExample02(c *check.C) {
	app := test.NewNeheExample02Application()

	app.Init(suite.OpenGl(), suite.Width(), suite.Height())

	app.Render()

	suite.ThenScreenShouldMatchReference(c, "NeheExample02.png")
}
