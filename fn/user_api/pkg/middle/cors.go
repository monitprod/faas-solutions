package middle

import (
	"github.com/aws/aws-lambda-go/events"
)

func CORS(r **events.APIGatewayProxyResponse) {

	// TODO: Improve this CORS
	(*r).Headers = map[string]string{
		"Access-Control-Allow-Origin":      "*",    // Required for CORS support to work
		"Access-Control-Allow-Credentials": "true", // Required for cookies, authorization headers with HTTPS
		"Access-Control-Allow-Headers":     "Content-Type",
		"Access-Control-Allow-Methods":     "GET, POST, PUT, DELETE, OPTIONS",
	}

}
