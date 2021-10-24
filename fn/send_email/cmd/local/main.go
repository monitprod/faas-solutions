package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	c "github.com/monitprod/send_email/pkg/constant"
	"github.com/monitprod/send_email/pkg/handler"
	"github.com/monitprod/send_email/pkg/util"
	f "github.com/monitprod/send_email/pkg/vo/function"
)

func main() {
	localFunc(nil)
}

func localFunc(payload *f.EventPayload) {
	ctx := context.WithValue(context.Background(), c.IsLocal, true)
	ctx = context.WithValue(ctx, c.LocalMainFunc, localFunc)

	if payload == nil {
		payload = startPayloadFromFile()
	}

	HandleRequest, err := handler.HandleRequest(ctx, *payload)

	if err != nil {
		log.Fatalln("Handle request failure:\n", err)
	}

	log.Printf("%+v", HandleRequest)
}

func startPayloadFromFile() *f.EventPayload {
	payloadFile := util.GetRootPath() + "/payload.json"

	file, err := ioutil.ReadFile(payloadFile)

	if err != nil {
		log.Fatalln("Payload read file error:\n", err)
	}

	payload := f.EventPayload{}

	err = json.Unmarshal([]byte(file), &payload)

	if err != nil {
		log.Fatalln("Failed to unmarshal payload:\n", err)
	}

	return &payload
}
