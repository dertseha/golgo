package app

const (
	DefaultWidth  int = 640
	DefaultHeight     = 480
)

const (
	DefaultName string = "Game of Life in Go"

	Version = "0.0.4"
)

type ApplicationParameter interface {
	Width() int
	Height() int

	Title() string
}

func DefaultApplicationParameter() ApplicationParameter {
	param := &basicApplicationParameter{
		width:  DefaultWidth,
		height: DefaultHeight,

		name:    DefaultName,
		version: Version}

	return param
}

type basicApplicationParameter struct {
	width  int
	height int

	name    string
	version string
}

func (param *basicApplicationParameter) Width() int {
	return param.width
}

func (param *basicApplicationParameter) Height() int {
	return param.height
}

func (param *basicApplicationParameter) Title() string {
	return param.name + " (v" + param.version + ")"
}
