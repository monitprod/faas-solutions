package util

import "os"

func GetRootPath() string {
	rootPath := os.Getenv("UA_ROOT_PATH")

	if rootPath == "" {
		panic("UA_ROOT_PATH environment variable was not started")
	}

	return rootPath
}
