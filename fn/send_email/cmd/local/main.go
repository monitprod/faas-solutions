package main

import (
	"github.com/monitprod/core/pkg/util/local"
	"github.com/monitprod/send_email/pkg/handler"
	"github.com/monitprod/send_email/pkg/util"
)

func main() {
	localFunc(nil)
}

func localFunc(payload *map[string]interface{}) {
	local.Start(local.StartBuilder{
		Handler:     handler.HandleRequest,
		LocalFunc:   localFunc,
		Payload:     payload,
		PayloadFile: util.GetRootPath() + "/payload.json",
	})
}
