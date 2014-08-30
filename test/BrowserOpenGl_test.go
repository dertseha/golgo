package test

import (
	"time"

	check "gopkg.in/check.v1"
)

type BrowserOpenGlExamplesSuite struct {
	BrowserSuite
}

var _ = check.Suite(&BrowserOpenGlExamplesSuite{})

func (suite *BrowserOpenGlExamplesSuite) TestNeheExample02(c *check.C) {
	suite.wd.Get("localhost:8080/golgo-js.html")

	time.Sleep(time.Millisecond * 100)

	suite.ThenScreenShouldMatchReference("NeheExample02.png")
}
