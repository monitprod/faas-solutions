package handler

import (
	"context"
	"fmt"

	f "github.com/monitprod/send_email/pkg/interface/function"
)

func HandleRequest(ctx context.Context, payload f.EventPayload) (f.Response, error) {
	return f.Response{
		Message: fmt.Sprintf("Lambda Started!\n%+v", payload),
	}, nil
}
