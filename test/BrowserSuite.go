package test

import (
	"image/png"
	"net"
	"net/http"

	"bitbucket.org/tebeka/selenium"

	testUtil "github.com/dertseha/golgo/test/util"

	check "gopkg.in/check.v1"
)

// http://godoc.org/bitbucket.org/tebeka/selenium

type BrowserSuite struct {
	listener net.Listener

	c  *check.C
	wd selenium.WebDriver
}

func (suite *BrowserSuite) SetUpSuite(c *check.C) {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("../"))))
	listener, err := net.Listen("tcp", ":8080")
	c.Assert(err, check.IsNil)
	go http.Serve(listener, nil)

	suite.listener = listener
}

func (suite *BrowserSuite) TearDownSuite(c *check.C) {
	suite.listener.Close()
}

func (suite *BrowserSuite) SetUpTest(c *check.C) {
	caps := selenium.Capabilities{"browserName": "firefox"}
	executor := ""

	wd, err := selenium.NewRemote(caps, executor)
	if err != nil {
		panic(err)
	}

	suite.c = c
	suite.wd = wd
}

func (suite *BrowserSuite) TearDownTest(c *check.C) {
	suite.wd.Quit()
}

func (suite *BrowserSuite) ThenScreenShouldMatchReference(refName string) {
	screenshotData, _ := suite.wd.Screenshot()
	liveImg, err := png.Decode(testUtil.ReadSlice(screenshotData))
	suite.c.Assert(err, check.IsNil)

	testUtil.SaveImage("screenshots/browser/"+refName, liveImg)
	confidence := testUtil.GetMatch("resources/"+refName, liveImg)

	suite.c.Check(confidence >= 0.99, check.Equals, true, check.Commentf("Confidence == %f", confidence))
}
