package route

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/monitprod/user_api/pkg/handler/product"
)

type Route struct {
	Method string
	Path   string
}

type HandleFuncGateway = func(context.Context, events.APIGatewayProxyRequest) (
	*events.APIGatewayProxyResponse, error,
)

var Routes = map[Route]HandleFuncGateway{
	// Use lower case!
	{Method: "get", Path: "/products"}: product.HandleProductRequest,
	{Method: "put", Path: "/sign"}:     product.HandleProductRequest,
}
