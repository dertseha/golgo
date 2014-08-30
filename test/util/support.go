package util

import "os"

func IsWebGlSupported() bool {
	isOnDroneIo := os.ExpandEnv("$DRONE") == "true"

	return !isOnDroneIo
}
