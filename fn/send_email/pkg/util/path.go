package util

import "os"

func GetRootPath() string {
	rootPath := os.Getenv("SE_ROOT_PATH")

	if rootPath == "" {
		panic("SE_ROOT_PATH environment variable was not started")
	}

	return rootPath
}
