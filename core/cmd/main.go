package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/monitprod/core"
	"github.com/monitprod/core/pkg/repository"
	u "github.com/monitprod/core/pkg/util"
)

func main() {
	fmt.Println("Core Started!\n" +
		"This execution is only for tests")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	core.StartRepository(ctx)

	userRepository := repository.NewUserRepositoryMongoDB()

	users, err := userRepository.GetUsers(ctx,
		repository.GetUsersOptions{
			Page: u.PaginateOptions{
				CurrentPage: 0,
				PageSize:    1,
			},
		})

	if err != nil {
		log.Fatalln("Error while get users from repository", err)
	}

	fmt.Println(*users)

	/*// Mongodb Client
	client := database.GetClient()

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
	}*/

}
