package app

import (
	"github.com/dertseha/golgo/runner"
)

const (
	DefaultWidth  int = 640
	DefaultHeight     = 480
)

const (
	DefaultName string = "Game of Life in Go"

	Version = "0.0.5"
)

func DefaultApplicationParameter() runner.ApplicationParameter {
	return runner.BasicApplicationParameter(DefaultWidth, DefaultHeight, DefaultName+" (v"+Version+")")
}
