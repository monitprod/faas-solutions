package service

import (
	"context"
	"log"

	m "github.com/monitprod/db_repository/pkg/models"
	"github.com/monitprod/db_repository/pkg/repository"
	r "github.com/monitprod/db_repository/pkg/repository"
	"github.com/monitprod/db_repository/pkg/util"
	f "github.com/monitprod/send_email/pkg/interface/function"
)

type UserService interface {
	GetUsers(ctx context.Context) (*[]m.User, error)
}

type UserServiceImp struct {
	UserRepository r.UserRepository
	Payload        f.EventPayload
}

func newUserServiceImp(
	userRepository r.UserRepository,
	payload f.EventPayload) UserService {

	return &UserServiceImp{
		UserRepository: userRepository,
		Payload:        payload,
	}
}

func (e *UserServiceImp) GetUsers(ctx context.Context) (*[]m.User, error) {

	users, err := e.UserRepository.GetUsers(ctx, repository.GetUsersOptions{
		Page: util.PaginateOptions{
			CurrentPage: e.Payload.Execution,
			PageSize:    e.Payload.UsersPerExecution,
		},
	})

	if err != nil {
		log.Fatalln("Error while get users from repository:\n", err)
		return nil, err
	}

	return users, nil
}
