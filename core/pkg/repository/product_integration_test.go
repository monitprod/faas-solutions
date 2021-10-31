// +build integration

package repository

import (
	"context"
	log "github.com/sirupsen/logrus"
	"testing"

	"github.com/monitprod/core"
	"github.com/monitprod/core/pkg/vo"
)

var getProductsOptionsMock = GetProductsOptions{
	Page: vo.PaginateOptions{
		CurrentPage: 1,
		PageSize:    1,
	},
}

func TestProductsRepository(t *testing.T) {

	core.UseCoreSmp(func(ctx context.Context) {

		productRepository := NewProductRepositoryMongoDB()

		products, err := productRepository.GetProducts(ctx,
			getProductsOptionsMock)

		if err != nil {
			log.Errorln("Error while get products from repository", err)
		}

		log.Println(*products)

	})

}
