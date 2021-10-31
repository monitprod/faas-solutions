package service

import (
	"context"
	"os"

	coreConstant "github.com/monitprod/core/pkg/constant"
	coreService "github.com/monitprod/core/pkg/service"
	"github.com/monitprod/core/pkg/util/local"
	f "github.com/monitprod/send_email/pkg/vo/function"
)

type SendEmailFnService interface {
	RunNewFunction(ctx context.Context, payload f.EventPayload) error
}

type SendEmailFnServiceImp struct {
}

func NewSendEmailFnServiceImp() SendEmailFnService {
	return &SendEmailFnServiceImp{}
}

func (e *SendEmailFnServiceImp) RunNewFunction(ctx context.Context, payload f.EventPayload) error {
	isLocal, _ := ctx.Value(coreConstant.IsLocal).(bool)
	localMainFunc, _ := ctx.Value(coreConstant.LocalMainFunc).(local.LocalFunc)

	builder := coreService.FunctionBuilder{
		IsLocal:   isLocal,
		LocalFunc: localMainFunc,
		Payload:   payload.ToMap(),
	}

	if !isLocal {
		builder.ServiceOptions = &coreService.ServiceOptions{
			Region:       os.Getenv("SE_FUNCTION_REGION"),
			FunctionName: os.Getenv("SE_FUNCTION_NAME"),
			Credentials: coreService.Credentials{
				AccessKey: os.Getenv("SE_AWS_ACCESS_KEY_ID"),
				Secret:    os.Getenv("SE_AWS_SECRET"),
			},
		}
	}

	funcService := coreService.NewFunctionServiceImp(builder)

	return funcService.Exec()
}
