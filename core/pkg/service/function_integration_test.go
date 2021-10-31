// +build integration

package service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"os"
	"testing"

	"github.com/monitprod/core"
)

func TestExecLambda(t *testing.T) {
	var payloadMock = map[string]interface{}{
		"execution":         1,
		"usersPerExecution": 1,
	}

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
			log.Errorln("Error while execute new FaaS function:", err)
		}
	})

}
