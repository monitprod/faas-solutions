package handler

import (
	"context"
	"fmt"

	"github.com/monitprod/core"

	"github.com/monitprod/send_email/pkg/util"
	f "github.com/monitprod/send_email/pkg/vo/function"
)

func HandleRequest(ctx context.Context, payload f.EventPayload) (f.Response, error) {
	util.StartEnv()

	core.UseCore(ctx, func() error {
		return sendEmailHandler(ctx, payload)
	})

	return f.Response{
		Message: fmt.Sprintf("Lambda Started!\n%+v", payload),
	}, nil
}
