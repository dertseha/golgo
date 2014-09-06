package util

import (
	"path"
	"runtime"
)

func GetProjectRootPath() string {
	_, localFileName, _, _ := runtime.Caller(0)

	return path.Join(path.Dir(localFileName), "..", "..")
}
