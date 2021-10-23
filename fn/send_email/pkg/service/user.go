package service

import (
	"context"
	"log"

	m "github.com/monitprod/core/pkg/models"
	r "github.com/monitprod/core/pkg/repository"
	"github.com/monitprod/core/pkg/vo"
	f "github.com/monitprod/send_email/pkg/vo/function"
)

type UserService interface {
	GetUsers(ctx context.Context) (*[]m.User, error)
	CountUsers(ctx context.Context) (*int64, error)
}

type UserServiceImp struct {
	UserRepository r.UserRepository
	Payload        f.EventPayload
}

func NewUserServiceImp(
	userRepository r.UserRepository,
	payload f.EventPayload) UserService {

	return &UserServiceImp{
		UserRepository: userRepository,
		Payload:        payload,
	}
}

func (e *UserServiceImp) GetUsers(ctx context.Context) (*[]m.User, error) {

	users, err := e.UserRepository.GetUsers(ctx, r.GetUsersOptions{
		Page: vo.PaginateOptions{
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

func (e *UserServiceImp) CountUsers(ctx context.Context) (*int64, error) {
	count, err := e.UserRepository.Count(ctx, true)

	if err != nil {
		log.Fatalln("Error while count users from repository:\n", err)
		return nil, err
	}

	return count, nil
}
