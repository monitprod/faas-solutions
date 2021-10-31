package service

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/monitprod/core/pkg/models"
	r "github.com/monitprod/core/pkg/repository"
)

type UserService interface {
	SignUser(ctx context.Context, user *models.User) error
}

type UserServiceImp struct {
	UserRepository r.UserRepository
}

func NewUserServiceImp(
	userRepository r.UserRepository,
) UserService {

	return &UserServiceImp{
		UserRepository: userRepository,
	}
}

func (u *UserServiceImp) SignUser(ctx context.Context, user *models.User) error {

	err := u.UserRepository.Create(ctx, user)

	if err != nil {
		log.Errorln("Error while sign user from repository:\n", err)
		return err
	}

	return nil
}
