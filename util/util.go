package util

import "runtime"

func IsCloudShell() bool {
	return runtime.GOOS == "js"
}