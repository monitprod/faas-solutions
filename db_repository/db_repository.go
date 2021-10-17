package db_repository

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/monitprod/db_repository/pkg/loaders"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	Client *mongo.Client
}

func StartRepository(ctx context.Context) Repository {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln("Error: Dot env not initialized")
	}

	return Repository{
		Client: loaders.MongoConnect(ctx),
	}
}
