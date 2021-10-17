package repository

import (
	"context"
	"log"
	"testing"

	"github.com/monitprod/db_repository"
	u "github.com/monitprod/db_repository/pkg/util"
)

func TestUsersRepository(t *testing.T) {

	ctx := context.Background()
	db_repository.StartRepository(ctx)

	log.Println("DB Repository Started!")

	userRepository := NewUserRepositoryMongoDB()

	users, err := userRepository.GetUsers(ctx,
		GetUsersOptions{
			Page: u.PaginateOptions{
				CurrentPage: 0,
				PageSize:    1,
			},
		})

	if err != nil {
		log.Fatalln("Error while get users from repository", err)
	}

	log.Println(*users)

}
