package repository

import (
	"context"
	"log"
	"testing"

	"github.com/monitprod/core"
	"github.com/monitprod/core/pkg/vo"
)

func TestUsersRepository(t *testing.T) {

	core.UseCoreSmp(func(ctx context.Context) {
		userRepository := NewUserRepositoryMongoDB()

		users, err := userRepository.GetUsers(ctx,
			GetUsersOptions{
				Page: vo.PaginateOptions{
					CurrentPage: 0,
					PageSize:    1,
				},
			})

		if err != nil {
			log.Fatalln("Error while get users from repository", err)
		}

		log.Println(*users)
	})
}
