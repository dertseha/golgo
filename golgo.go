package main

import (
	"github.com/dertseha/golgo/app"
	runner "github.com/dertseha/golgo/native/glfw"
)

func main() {
	param := app.DefaultApplicationParameter()
	application := app.NewMainApplication()

	runner.Run(application, param)
}
