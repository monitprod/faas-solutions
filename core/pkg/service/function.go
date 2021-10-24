package service

import (
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/monitprod/send_email/pkg/vo/function"
)

// FunctionService is a service to execute new FaaS Function
type FunctionService interface {
	Exec() error
}

type FunctionBuilder struct {
	Payload map[string]interface{}
	IsLocal bool

	LocalFunc func(payload *function.EventPayload)

	// ServiceOptions can be nil if execution is locally
	ServiceOptions *ServiceOptions
}

type ServiceOptions struct {
	Region       string
	FunctionName string
}

type FunctionServiceImp struct {
	Builder FunctionBuilder
}

func NewFunctionServiceImp(builder FunctionBuilder) FunctionService {
	return &FunctionServiceImp{builder}
}

func (f FunctionServiceImp) Exec() error {
	if f.Builder.IsLocal {
		return f.localExec()
	}

	return f.lambdaExec()
}

func (f FunctionServiceImp) localExec() error {
	log.Println("Start local function execution")

	payload, err := function.EventPayloadFromMap(f.Builder.Payload)

	if err != nil {
		return err
	}

	f.Builder.LocalFunc(payload)

	return nil
}

func (f FunctionServiceImp) lambdaExec() error {
	log.Println("Start lambda function execution")

	sess := session.Must(
		session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}),
	)

	serviceOptions := f.Builder.ServiceOptions

	config := aws.Config{
		Region: aws.String(serviceOptions.Region),
	}

	client := lambda.New(sess, &config)

	request := f.Builder.Payload

	payload, err := json.Marshal(request)
	if err != nil {
		log.Println("Error marshalling MyGetItemsFunction request")
		os.Exit(0)
	}

	invokeInput := lambda.InvokeInput{
		FunctionName:   aws.String(serviceOptions.FunctionName),
		Payload:        payload,
		InvocationType: aws.String("Event"), // Async
	}

	// TODO: To retain events that were not processed, configure your
	// function with a dead-letter queue
	// (https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#dlq).

	_, err = client.Invoke(&invokeInput)

	if err != nil {
		log.Println("Error calling MyGetItemsFunction")
		os.Exit(0)
	}

	return nil
}
