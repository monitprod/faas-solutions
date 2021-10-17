package repository

import (
	"context"
	"log"

	c "github.com/monitprod/db_repository/pkg/constant"
	"github.com/monitprod/db_repository/pkg/loaders/database"
	m "github.com/monitprod/db_repository/pkg/models"
	"github.com/monitprod/db_repository/pkg/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GetUsersOptions struct {
	Page util.PaginateOptions
}

type UserRepository interface {
	GetUsers(ctx context.Context, opt GetUsersOptions) (*[]m.User, error)
}

type UserRepositoryMongoDB struct {
}

func NewUserRepositoryMongoDB() UserRepository {
	return UserRepositoryMongoDB{}
}

func (u UserRepositoryMongoDB) GetUsers(ctx context.Context, opt GetUsersOptions) (*[]m.User, error) {
	// Mongodb Client
	client := database.GetClient()

	userCollection := client.
		Database(c.Database).
		Collection(c.UserCollection)

	var users []m.User

	findOptions := options.FindOptions{}

	cursor, err := userCollection.Find(
		ctx,
		bson.M{},
		util.PaginateFind(&findOptions, opt.Page),
	)

	if err != nil {
		log.Fatalln("Error while find user collection:", err)
	}

	if err = cursor.All(ctx, &users); err != nil {
		log.Fatalln("Error while cursor all users of user collection:", err)
	}

	return &users, nil
}
