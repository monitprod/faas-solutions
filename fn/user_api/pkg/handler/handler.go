package handler

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/monitprod/core"
	"github.com/monitprod/user_api/pkg/handler/route"
	"github.com/monitprod/user_api/pkg/util"
)

func HandleRequest(ctx context.Context, payload map[string]interface{}) (map[string]interface{}, error) {
	util.StartEnv()

	req, err := requestFromMap(payload)
	if err != nil {
		return nil, err
	}

	var res *events.APIGatewayProxyResponse
	err = core.UseCore(ctx, func() (err error) {
		res, err = route.HandleAPIGatewayRoutes(ctx, *req)
		return err
	})

	return responseToMap(res), err
}

func responseToMap(e *events.APIGatewayProxyResponse) (res map[string]interface{}) {
	a, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(a, &res)
	return
}

func requestFromMap(m map[string]interface{}) (*events.APIGatewayProxyRequest, error) {
	r := events.APIGatewayProxyRequest{}

	data, err := json.Marshal(m)
	if err == nil {
		err = json.Unmarshal(data, &r)
	}
	return &r, err
}
