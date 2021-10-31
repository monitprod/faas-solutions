package service

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/monitprod/core/pkg/util/local"
)

// FunctionService is a service to execute new FaaS Function
type FunctionService interface {
	Exec() error
}

type FunctionBuilder struct {
	Payload map[string]interface{}
	IsLocal bool

	LocalFunc local.LocalFunc

	// ServiceOptions can be nil if execution is locally
	ServiceOptions *ServiceOptions
}

type ServiceOptions struct {
	Region       string
	FunctionName string
	Credentials  Credentials
}

type Credentials struct {
	AccessKey string
	Secret    string
	Token     string // Is Optional
}

type FunctionServiceImp struct {
	Builder FunctionBuilder
}

func NewFunctionServiceImp(builder FunctionBuilder) FunctionService {
	return &FunctionServiceImp{builder}
}

// TODO: Document this
func (f FunctionServiceImp) Exec() error {
	if f.Builder.IsLocal {
		return f.localExec()
	}

	return f.lambdaExec()
}

func (f FunctionServiceImp) localExec() error {
	log.Println("Start local function execution")

	f.Builder.LocalFunc(&f.Builder.Payload)

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

	cred := f.Builder.ServiceOptions.Credentials

	config := aws.Config{
		Credentials: credentials.NewStaticCredentials(
			cred.AccessKey,
			cred.Secret,
			cred.Token,
		),
		Region: aws.String(serviceOptions.Region),
	}

	client := lambda.New(sess, &config)

	request := f.Builder.Payload

	payload, err := json.Marshal(request)
	if err != nil {
		log.Errorln("Error marshalling function request:", err)
		return err
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
		log.Errorln("Error calling function:", err)
		return err
	}

	return nil
}
