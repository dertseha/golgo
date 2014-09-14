package runner

type ApplicationParameter interface {
	Width() int
	Height() int

	Title() string
}
