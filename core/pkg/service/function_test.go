package service

import (
	"log"
	"testing"

	"github.com/monitprod/core/pkg/util/local"
)

func TestExecLocal(t *testing.T) {

	var payloadMock = map[string]interface{}{
		"execution":         1,
		"usersPerExecution": 1,
	}

	var localFuncMock local.LocalFunc = func(payload *map[string]interface{}) {
		// My func. . .
	}

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
