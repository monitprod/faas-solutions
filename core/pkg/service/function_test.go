package service

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/monitprod/core"
)

var payloadMock = map[string]interface{}{
	"execution":         1,
	"usersPerExecution": 1,
}

var localFuncMock = func(payload map[string]interface{}) {
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
			},
		}

		funcService := NewFunctionServiceImp(builder)

		err := funcService.Exec()
		if err != nil {
			log.Fatalln("Error while execute new FaaS function:", err)
		}
	})

}
