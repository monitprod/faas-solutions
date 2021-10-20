package handler

import (
	"context"
	"fmt"
	"log"

	"github.com/monitprod/core"
	"github.com/monitprod/core/pkg/loaders/database"
	f "github.com/monitprod/send_email/pkg/interface/function"
)

func HandleRequest(ctx context.Context, payload f.EventPayload) (f.Response, error) {
	core.UseCore(ctx, func() error {
		a := database.GetClient().NumberSessionsInProgress()
		log.Print(a)
		return nil
	})

	return f.Response{
		Message: fmt.Sprintf("Lambda Started!\n%+v", payload),
	}, nil
}
