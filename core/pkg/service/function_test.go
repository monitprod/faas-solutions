package service

import (
	"log"
	"testing"
)

func TestExec(t *testing.T) {

	var builder FunctionBuilder = FunctionBuilder{
		Payload:        map[string]interface{}{"teste": "1"},
		IsLocal:        true,
		ServiceOptions: nil,
	}

	funcService := NewFunctionServiceImp(builder)

	err := funcService.Exec()
	if err != nil {
		log.Fatalln("Error while execute new FaaS function:", err)
	}

}
