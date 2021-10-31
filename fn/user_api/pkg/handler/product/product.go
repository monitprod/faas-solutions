package product

import (
	"context"
	"encoding/json"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-lambda-go/events"
	"github.com/monitprod/core/pkg/models"
	r "github.com/monitprod/core/pkg/repository"

	s "github.com/monitprod/user_api/pkg/service"
)

type ProductsResponse struct {
	Products *[]models.Product `json:"products"`
}

func HandleProductRequest(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var products *[]models.Product

	products, err := getProductsHandler(ctx, request)
	if err != nil {
		return nil, err
	}

	res := ProductsResponse{
		Products: products,
	}

	resJson, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}

	return &events.APIGatewayProxyResponse{
		Body:       string(resJson),
		StatusCode: 200,
	}, nil
}

func getProductsHandler(ctx context.Context, request events.APIGatewayProxyRequest) (*[]models.Product, error) {

	productRepo := r.NewProductRepositoryMongoDB()
	productService := s.NewProductServiceImp(productRepo)

	products, err := productService.GetProducts(ctx)
	if err != nil {
		log.Errorln("Error while get products from product service", err)
		return nil, err
	}

	return products, nil
}
