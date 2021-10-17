package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func ConnectClient(ctx context.Context) *mongo.Client {

	uri := os.Getenv("DB_REPOSITORY_MONGO_URI")

	clientOptions := options.Client().ApplyURI(uri)

	var err error
	client, err = mongo.NewClient(clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	return client
}

func GetClient() *mongo.Client {
	if isClientStarted() {
		return client
	}

	return nil
}

func DisconnectClient(ctx context.Context) error {
	if isClientStarted() {
		return client.Disconnect(ctx)
	}

	return nil
}

func isClientStarted() bool {
	if client == nil {
		log.Fatalln("Mongo Client not started, use db_repository.StartRepository(ctx) to start")
		return false
	}
	return true
}
