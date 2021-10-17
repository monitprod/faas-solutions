package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/monitprod/db_repository"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	fmt.Println("DB Repository Started!\n" +
		"This execution is only for tests")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	repository := db_repository.StartRepository(ctx)

	// Mongodb Client
	client := repository.Client

	productCollection := client.Database("monitprod").Collection("products")

	ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)

	cur, err := productCollection.Find(ctx, bson.D{})

	if err != nil {
		log.Fatalln("Error while find product collection")
	}

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("%+v", result)

	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

}
