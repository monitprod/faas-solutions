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

// UseCoreSmp is simple form of UseCore method
func UseCoreSmp(execution func(ctx context.Context)) {
	ctx := context.Background()

	UseCore(ctx, func() error {
		execution(ctx)
		return nil
	})
}

func UseCore(ctx context.Context, execution func() error) error {
	start(ctx)

	err := execution()

	defer close(ctx)

	if err != nil {
		log.Fatalln("Error at execution on UseCore\n", err)

		return err
	}

	return nil
}

func start(ctx context.Context) {
	startRootPath()

	err := godotenv.Load(getDefaultEnvPath())

	if err != nil {
		log.Println("INFO: Core dot env not initialized:", err)
	}

	database.ConnectClient(ctx)

	log.Println("Core Started!")
}

func close(ctx context.Context) {
	database.DisconnectClient(ctx)
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
