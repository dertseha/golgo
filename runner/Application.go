package runner

import (
	gles "github.com/dertseha/golgo/gles2"
)

type Application interface {
	Init(gl gles.OpenGl, width, height int)
	Resize(width, height int)
	Render()
}
