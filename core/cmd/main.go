package main

import (
	"context"
	"fmt"
	"log"

	"github.com/monitprod/core"
	"github.com/monitprod/core/pkg/repository"
	"github.com/monitprod/core/pkg/vo"
)

func main() {
	fmt.Println("This execution is only for tests")

	core.UseCoreSmp(func(ctx context.Context) {
		userRepository := repository.NewUserRepositoryMongoDB()

		users, err := userRepository.GetUsers(ctx,
			repository.GetUsersOptions{
				Page: vo.PaginateOptions{
					CurrentPage: 0,
					PageSize:    1,
				},
			})

		if err != nil {
			log.Fatalln("Error while get users from repository", err)
		}

		fmt.Println(*users)
	})
}
