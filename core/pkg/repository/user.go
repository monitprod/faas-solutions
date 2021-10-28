package repository

import (
	"context"
	"log"

	c "github.com/monitprod/core/pkg/constant"
	"github.com/monitprod/core/pkg/loaders/database"
	m "github.com/monitprod/core/pkg/models"
	"github.com/monitprod/core/pkg/util"
	"github.com/monitprod/core/pkg/vo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GetUsersOptions struct {
	Page vo.PaginateOptions
}

type UserRepository interface {
	GetUsers(ctx context.Context, opt GetUsersOptions) (*[]m.User, error)
	Count(ctx context.Context, estimated bool) (*int64, error)
	Create(ctx context.Context, user *m.User) error
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
		return nil, err
	}

	if err = cursor.All(ctx, &users); err != nil {
		log.Fatalln("Error while cursor all users of user collection:", err)
		return nil, err
	}

	return &users, nil
}

// Count - "estimated" is more performatic, but less accurate
func (u UserRepositoryMongoDB) Count(ctx context.Context, estimated bool) (*int64, error) {
	// Mongodb Client
	client := database.GetClient()

	userCollection := client.
		Database(c.Database).
		Collection(c.UserCollection)

	var count int64
	var err error

	if estimated {
		count, err = userCollection.EstimatedDocumentCount(ctx)
	} else {
		count, err = userCollection.CountDocuments(ctx, bson.D{})
	}

	if err != nil {
		log.Fatalln("Error while count user collection:", err)
	}

	return &count, nil
}

func (u UserRepositoryMongoDB) Create(ctx context.Context, user *m.User) error {
	// Mongodb Client
	client := database.GetClient()

	userCollection := client.
		Database(c.Database).
		Collection(c.UserCollection)

	_, err := userCollection.InsertOne(ctx, *user)

	if err != nil {
		log.Fatalln("Error while create user:", err)
		return err
	}

	return nil
}
