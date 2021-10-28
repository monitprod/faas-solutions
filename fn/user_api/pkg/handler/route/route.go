package route

import (
	"context"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/monitprod/user_api/pkg/route"
)

func HandleAPIGatewayRoutes(ctx context.Context, request events.APIGatewayProxyRequest) (
	events.APIGatewayProxyResponse, error,
) {

	key := route.Route{
		Method: strings.ToLower(request.HTTPMethod),
		Path:   strings.ToLower(request.Path),
	}

	return route.Routes[key](ctx, request)
}
