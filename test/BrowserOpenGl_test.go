package test

import (
	"time"

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
	suite.wd.Get("localhost:8080/golgo-js.html")

	time.Sleep(time.Millisecond * 100)

	suite.ThenScreenShouldMatchReference(c, "NeheExample02.png")
}
