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

type GetProductsOptions struct {
	Page util.PaginateOptions
}

type ProductRepository interface {
	GetProducts(ctx context.Context, opt GetProductsOptions) (*[]m.Product, error)
}

type ProductRepositoryMongoDB struct {
}

func NewProductRepositoryMongoDB() ProductRepository {
	return ProductRepositoryMongoDB{}
}

func (u ProductRepositoryMongoDB) GetProducts(ctx context.Context, opt GetProductsOptions) (*[]m.Product, error) {
	// Mongodb Client
	client := database.GetClient()

	productCollection := client.
		Database(c.Database).
		Collection(c.ProductCollection)

	var products []m.Product

	findOptions := options.FindOptions{}

	cursor, err := productCollection.Find(
		ctx,
		bson.M{},
		util.PaginateFind(&findOptions, opt.Page),
	)

	if err != nil {
		log.Fatalln("Error while find product collection:", err)
	}

	if err = cursor.All(ctx, &products); err != nil {
		log.Fatalln("Error while cursor all products of product collection:", err)
	}

	return &products, nil
}
