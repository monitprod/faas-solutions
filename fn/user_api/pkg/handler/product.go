package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/monitprod/core/pkg/models"
	r "github.com/monitprod/core/pkg/repository"

	s "github.com/monitprod/user_api/pkg/service"
)

func handleProductRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var products *[]models.Product

	products, err := getProductsHandler(ctx, request)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 500,
		}, nil
	}

	res := map[string]interface{}{
		"products": *products,
	}

	resJsonB, err := json.Marshal(res)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 500,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("%s", resJsonB),
		StatusCode: 200,
	}, nil
}

func getProductsHandler(ctx context.Context, request events.APIGatewayProxyRequest) (*[]models.Product, error) {

	productRepo := r.NewProductRepositoryMongoDB()
	productService := s.NewProductServiceImp(productRepo)

	products, err := productService.GetProducts(ctx)
	if err != nil {
		log.Fatalln("Error while get products from product service", err)
		return nil, err
	}

	return products, nil
}
