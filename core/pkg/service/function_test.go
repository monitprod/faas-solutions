package service

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/monitprod/core"
	"github.com/monitprod/core/pkg/util/local"
)

var payloadMock = map[string]interface{}{
	"execution":         1,
	"usersPerExecution": 1,
}

var localFuncMock local.LocalFunc = func(payload *map[string]interface{}) {
	// My func. . .
}

func TestExecLocal(t *testing.T) {

	var builder FunctionBuilder = FunctionBuilder{
		Payload:        payloadMock,
		IsLocal:        true,
		LocalFunc:      localFuncMock,
		ServiceOptions: nil,
	}

	funcService := NewFunctionServiceImp(builder)

	err := funcService.Exec()
	if err != nil {
		log.Fatalln("Error while execute new FaaS function:", err)
	}

}

func TestExecLambda(t *testing.T) {

	core.UseCoreSmp(func(ctx context.Context) {
		var builder FunctionBuilder = FunctionBuilder{
			Payload: payloadMock,
			IsLocal: false,
			ServiceOptions: &ServiceOptions{
				Region:       os.Getenv("CORE_AWS_REGION"),
				FunctionName: os.Getenv("CORE_AWS_FUNCTION_NAME"),
				Credentials: Credentials{
					AccessKey: os.Getenv("CORE_AWS_ACCESS_KEY_ID"),
					Secret:    os.Getenv("CORE_AWS_SECRET"),
					Token:     os.Getenv("CORE_AWS_TOKEN"),
				},
			},
		}

		funcService := NewFunctionServiceImp(builder)

		err := funcService.Exec()
		if err != nil {
			log.Fatalln("Error while execute new FaaS function:", err)
		}
	})

}
