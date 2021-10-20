package core

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/monitprod/core/pkg/loaders/database"
)

func StartRepositoryEnv(ctx context.Context, envPath string) {

	if envPath == "" {
		envPath = getDefaultEnvPath()
	}

	err := godotenv.Load(envPath)

	if err != nil {
		log.Fatalln("Error: Dot env not initialized:", err)
	}

	database.ConnectClient(ctx)
}

func StartRepository(ctx context.Context) {
	startRootPath()

	err := godotenv.Load(getDefaultEnvPath())

	if err != nil {
		log.Fatalln("Error: Dot env not initialized:", err)
	}

	database.ConnectClient(ctx)
}

func getDefaultEnvPath() string {
	return os.Getenv("CORE_ROOT_PATH") + "\\.env"
}

func startRootPath() {
	var (
		_, b, _, _ = runtime.Caller(0)
		basepath   = filepath.Dir(b)
	)

	os.Setenv("CORE_ROOT_PATH", basepath)
}
