package util

import (
	"os"
)

func GetRootPath() string {
	rootPath := os.Getenv("CORE_ROOT_PATH")

	return rootPath
}
