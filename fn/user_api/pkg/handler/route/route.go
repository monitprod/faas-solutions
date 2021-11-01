package route

import (
	"context"
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"

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
		return nil, fmt.Errorf("route not found, expected:\n%+v", key)
	}
	// TODO: Adding accept json header method options

	// Route
	r, err := routeHandler(ctx, request)

	// Middle After Route
	middle.ResponseError(&r, err)
	middle.CORS(&r) // TODO: improve security of it

	log.Printf("Response: %+v\n", *r)

	return r, err
}
