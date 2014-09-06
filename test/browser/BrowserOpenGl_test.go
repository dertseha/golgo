package browser

import (
	testUtil "github.com/dertseha/golgo/test/util"

	check "gopkg.in/check.v1"
)

type BrowserOpenGlExamplesSuite struct {
	BrowserSuite
}

var _ = check.Suite(&BrowserOpenGlExamplesSuite{})

func (suite *BrowserOpenGlExamplesSuite) SetUpTest(c *check.C) {
	if testUtil.IsWebGlSupported() {
		suite.BrowserSuite.SetUpTest(c)
	} else {
		c.Skip("Skipping WebGL tests due to lack of support")
	}
}

func (suite *BrowserOpenGlExamplesSuite) TestNeheExample02(c *check.C) {
	suite.GivenARunningApplication(c, "github.com/dertseha/golgo/test", "test.NewNeheExample02Application()")

	suite.WhenTheApplicationWasRenderedOnce()

	suite.ThenScreenShouldMatchReference(c, "NeheExample02.png")
}

func (suite *BrowserOpenGlExamplesSuite) TestClearBackground(c *check.C) {
	suite.GivenARunningApplication(c, "github.com/dertseha/golgo/test", "test.NewClearBackgroundApplication()")

	suite.WhenTheApplicationWasRenderedOnce()

	suite.ThenScreenShouldMatchReference(c, "ClearBackground.png")
}
