package loaders

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnect(ctx context.Context) *mongo.Client {

	uri := os.Getenv("DB_REPOSITORY_MONGO_URI")

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.NewClient(clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	return client
}
