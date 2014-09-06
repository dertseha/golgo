package browser

import (
	"image/png"
	"net"
	"net/http"
	"time"

	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path"

	"bitbucket.org/tebeka/selenium"

	testUtil "github.com/dertseha/golgo/test/util"

	check "gopkg.in/check.v1"
)

const mainSource = `package main

import (
	"github.com/dertseha/golgo/app"
	runner "github.com/dertseha/golgo/golgo-js/runner"
	"%s"
)

func main() {
	param := app.DefaultApplicationParameter()
	application := %s

	runner.Run(application, param)
}
`

type BrowserSuite struct {
	listener net.Listener

	wd selenium.WebDriver
}

func (suite *BrowserSuite) SetUpSuite(c *check.C) {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(testUtil.GetProjectRootPath()))))
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

	suite.wd = wd
}

func (suite *BrowserSuite) TearDownTest(c *check.C) {
	suite.wd.Quit()

	os.Remove(suite.getMainGoFilePath())
}

func (suite *BrowserSuite) getMainGoFilePath() string {
	return path.Join(testUtil.GetProjectRootPath(), "build", "test.go")
}

func (suite *BrowserOpenGlExamplesSuite) GivenARunningApplication(c *check.C, importPath string, appConstructor string) {
	mainFile := suite.getMainGoFilePath()
	file, err := os.Create(mainFile)
	c.Assert(err, check.IsNil)
	writer := bufio.NewWriter(file)

	fmt.Fprintf(writer, mainSource, importPath, appConstructor)
	writer.Flush()
	file.Close()

	suite.GivenARunningMainFile(c, mainFile)
}

func (suite *BrowserOpenGlExamplesSuite) GivenARunningMainFile(c *check.C, mainFile string) {
	buildFileName := path.Join(testUtil.GetProjectRootPath(), "build", "test.js")

	bin := os.ExpandEnv("$GOPATH/bin/gopherjs")
	cmd := exec.Command(bin, "build", "-o", buildFileName, mainFile)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	result := cmd.Run()

	c.Assert(result, check.IsNil)

	suite.wd.Get("localhost:8080/test/golgo-test.html")
}

func (suite *BrowserSuite) WhenTheApplicationWasRenderedOnce() {
	time.Sleep(time.Millisecond * 100)
}

func (suite *BrowserSuite) ThenScreenShouldMatchReference(c *check.C, refName string) {
	screenshotPath := path.Join(testUtil.GetProjectRootPath(), "test", "screenshots", "browser", c.TestName()+"."+refName)
	referencePath := path.Join(testUtil.GetProjectRootPath(), "test", "resources", refName)
	screenshotData, _ := suite.wd.Screenshot()
	liveImg, err := png.Decode(testUtil.ReadSlice(screenshotData))
	c.Assert(err, check.IsNil)

	testUtil.SaveImage(screenshotPath, liveImg)
	confidence := testUtil.GetMatch(referencePath, liveImg)

	c.Check(confidence >= 0.99, check.Equals, true, check.Commentf("Confidence == %f", confidence))
}
