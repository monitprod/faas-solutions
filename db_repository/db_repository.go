package db_repository

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/monitprod/db_repository/pkg/loaders/database"
)

func StartRepository(ctx context.Context) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln("Error: Dot env not initialized")
	}

	database.ConnectClient(ctx)
}
