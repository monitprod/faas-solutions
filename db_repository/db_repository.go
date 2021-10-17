package db_repository

import (
	"context"
	"log"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/monitprod/db_repository/pkg/loaders/database"
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

	err := godotenv.Load(getDefaultEnvPath())

	if err != nil {
		log.Fatalln("Error: Dot env not initialized:", err)
	}

	database.ConnectClient(ctx)
}

func getDefaultEnvPath() string {
	return getBasepath() + "\\.env"
}

func getBasepath() string {
	var (
		_, b, _, _ = runtime.Caller(0)
		basepath   = filepath.Dir(b)
	)
	return basepath
}
