package route

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/monitprod/user_api/pkg/middle"
	"github.com/monitprod/user_api/pkg/route"
)

func HandleAPIGatewayRoutes(ctx context.Context, request events.APIGatewayProxyRequest) (
	*events.APIGatewayProxyResponse, error,
) {

	key := route.Route{
		Method: strings.ToLower(request.HTTPMethod),
		Path:   strings.ToLower(request.Path),
	}

	routeHandler, ok := route.Routes[key]
	if !ok {
		return nil, errors.New("route not found")
	}

	r, err := routeHandler(ctx, request)

	// Middle After Route
	middle.ResponseError(&r, err)

	log.Printf("Response: %+v\n", *r)

	return r, err
}
