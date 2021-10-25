package handler

import (
	"context"
	"fmt"

	"github.com/monitprod/core"

	"github.com/monitprod/send_email/pkg/util"
	f "github.com/monitprod/send_email/pkg/vo/function"
)

func HandleRequest(ctx context.Context, payload map[string]interface{}) (map[string]interface{}, error) {
	util.StartEnv()

	core.UseCore(ctx, func() error {
		p, err := f.EventPayloadFromMap(payload)
		if err != nil {
			return err
		}

		return sendEmailHandler(ctx, *p)
	})

	res := f.Response{
		Message: fmt.Sprintf("Lambda Started!\n%+v", payload),
	}

	return res.ToMap(), nil
}
