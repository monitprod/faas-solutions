package repository

import (
	"context"
	"log"
	"testing"

	"github.com/monitprod/db_repository"
	u "github.com/monitprod/db_repository/pkg/util"
)

func TestProductsRepository(t *testing.T) {

	ctx := context.Background()
	db_repository.StartRepository(ctx)

	log.Println("DB Repository Started!")

	productRepository := NewProductRepositoryMongoDB()

	products, err := productRepository.GetProducts(ctx,
		GetProductsOptions{
			Page: u.PaginateOptions{
				CurrentPage: 0,
				PageSize:    1,
			},
		})

	if err != nil {
		log.Fatalln("Error while get products from repository", err)
	}

	log.Println(*products)

}
