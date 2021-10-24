package service

import (
	"context"
	"log"

	m "github.com/monitprod/core/pkg/models"
	r "github.com/monitprod/core/pkg/repository"
	"github.com/monitprod/core/pkg/vo"
	f "github.com/monitprod/send_email/pkg/vo/function"
)

type ProductService interface {
	GetProducts(ctx context.Context) (*[]m.Product, error)
	CountProducts(ctx context.Context) (*int64, error)
}

type ProductServiceImp struct {
	ProductRepository r.ProductRepository
	Payload           f.EventPayload
}

func NewProductServiceImp(
	productRepository r.ProductRepository,
	payload f.EventPayload) ProductService {

	return &ProductServiceImp{
		ProductRepository: productRepository,
		Payload:           payload,
	}
}

func (e *ProductServiceImp) GetProducts(ctx context.Context) (*[]m.Product, error) {
	const limitProduct = 10

	products, err := e.ProductRepository.GetProducts(ctx, r.GetProductsOptions{
		Page: vo.PaginateOptions{
			CurrentPage: 1,
			PageSize:    limitProduct,
		},
	})

	if err != nil {
		log.Fatalln("Error while get products from repository:\n", err)
		return nil, err
	}

	return products, nil
}

func (e *ProductServiceImp) CountProducts(ctx context.Context) (*int64, error) {
	count, err := e.ProductRepository.Count(ctx, true)

	if err != nil {
		log.Fatalln("Error while count products from repository:\n", err)
		return nil, err
	}

	return count, nil
}
