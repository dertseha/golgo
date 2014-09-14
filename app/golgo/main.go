package main

import (
	"github.com/dertseha/golgo/app"
	"github.com/dertseha/golgo/core"
	"github.com/dertseha/golgo/runner/native"
)

func main() {
	param := app.DefaultApplicationParameter()
	application := core.NewMainApplication()

	native.Run(application, param)
}
