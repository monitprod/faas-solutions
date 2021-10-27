package handler

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"

	"github.com/monitprod/core"
	"github.com/monitprod/user_api/pkg/util"
)

func HandleRequest(ctx context.Context, payload map[string]interface{}) (map[string]interface{}, error) {
	util.StartEnv()

	var res events.APIGatewayProxyResponse

	req, err := RequestFromMap(payload)
	if err != nil {
		return nil, err
	}

	err = core.UseCore(ctx, func() (err error) {
		res, err = handleAPIGatewayRequest(ctx, *req)
		return err
	})

	return ResponseToMap(&res), err
}

func handleAPIGatewayRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return handleProductRequest(ctx, request)
}

func ResponseToMap(e *events.APIGatewayProxyResponse) (res map[string]interface{}) {
	a, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(a, &res)
	return
}

func RequestFromMap(m map[string]interface{}) (*events.APIGatewayProxyRequest, error) {
	r := events.APIGatewayProxyRequest{}

	data, err := json.Marshal(m)
	if err == nil {
		err = json.Unmarshal(data, &r)
	}
	return &r, err
}
