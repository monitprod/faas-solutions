package main

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/monitprod/core"
	"github.com/monitprod/core/pkg/repository"
	"github.com/monitprod/core/pkg/vo"
)

func main() {
	log.Println("This execution is only for tests")

	core.UseCoreSmp(func(ctx context.Context) {
		userRepository := repository.NewUserRepositoryMongoDB()

		users, err := userRepository.GetUsers(ctx,
			repository.GetUsersOptions{
				Page: vo.PaginateOptions{
					CurrentPage: 1,
					PageSize:    1,
				},
			})

		if err != nil {
			log.Errorln("Error while get users from repository", err)
		}

		log.Println(*users)
	})
}
