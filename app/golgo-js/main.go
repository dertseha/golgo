package main

import (
	"github.com/dertseha/golgo/app"
	"github.com/dertseha/golgo/core"
	"github.com/dertseha/golgo/runner/browser"
)

func main() {
	param := app.DefaultApplicationParameter()
	application := core.NewMainApplication()

	browser.Run(application, param)
}
