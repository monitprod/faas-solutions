package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/monitprod/user_api/pkg/handler"
)

func main() {
	lambda.Start(handler.HandleRequest)
}
