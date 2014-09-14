package runner

type basicApplicationParameter struct {
	width  int
	height int

	title string
}

func BasicApplicationParameter(width, height int, title string) ApplicationParameter {
	param := &basicApplicationParameter{
		width:  width,
		height: height,

		title: title}

	return param
}

func (param *basicApplicationParameter) Width() int {
	return param.width
}

func (param *basicApplicationParameter) Height() int {
	return param.height
}

func (param *basicApplicationParameter) Title() string {
	return param.title
}
