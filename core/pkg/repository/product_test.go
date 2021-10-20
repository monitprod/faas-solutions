package repository

import (
	"context"
	"log"
	"testing"

	"github.com/monitprod/core"
	u "github.com/monitprod/core/pkg/util"
)

func TestProductsRepository(t *testing.T) {

	ctx := context.Background()
	core.StartRepository(ctx)

	log.Println("Core Started!")

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
