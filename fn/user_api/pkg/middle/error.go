package middle

import (
	"github.com/aws/aws-lambda-go/events"
)

func ResponseError(r **events.APIGatewayProxyResponse, err error) {

	if err != nil {
		*r = &events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 500,
		}
	}
}
