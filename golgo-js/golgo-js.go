package main

import (
	"github.com/dertseha/golgo/app"
	runner "github.com/dertseha/golgo/golgo-js/runner"
)

func main() {
	param := app.DefaultApplicationParameter()
	application := app.NewMainApplication()

	runner.Run(application, param)
}
