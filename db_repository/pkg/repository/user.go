package repository

import (
	m "github.com/monitprod/db_repository/pkg/models"
)

type UserRepository interface {
	GetUsers() (m.User, error)
}

type UserRepositoryMongoDB struct {
}

func NewUserRepositoryMongoDB() UserRepository {
	return UserRepositoryMongoDB{}
}

func (u UserRepositoryMongoDB) GetUsers() (m.User, error) {
	return m.User{}, nil
}
