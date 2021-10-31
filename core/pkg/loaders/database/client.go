package database

import (
	"context"
	"errors"
	"os"

	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func ConnectClient(ctx context.Context) *mongo.Client {

	uri := os.Getenv("CORE_MONGO_URI")

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
	isStarted, err := isClientStarted()
	if err != nil {
		log.Errorln(err)
		return nil
	}

	if isStarted {
		return client
	}

	return nil
}

func DisconnectClient(ctx context.Context) error {
	isStarted, err := isClientStarted()
	if err != nil {
		return err
	}

	if isStarted {
		err := client.Disconnect(ctx)
		client = nil
		return err
	}

	return errors.New("DisconnectClient unrecognized error")
}

func isClientStarted() (bool, error) {
	if client == nil {
		return false, errors.New("mongo client not started, use core.StartRepository(ctx) to start")
	}
	return true, nil
}
