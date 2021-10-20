package repository

import (
	"context"
	"log"
	"testing"

	"github.com/monitprod/core"
	u "github.com/monitprod/core/pkg/util"
)

func TestUsersRepository(t *testing.T) {

	ctx := context.Background()
	core.StartRepository(ctx)

	log.Println("Core Started!")

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
