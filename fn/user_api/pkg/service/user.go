package service

import (
	r "github.com/monitprod/core/pkg/repository"
	f "github.com/monitprod/user_api/pkg/vo/function"
)

type UserService interface {
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
