package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/monitprod/send_email/pkg/handler"
	f "github.com/monitprod/send_email/pkg/interface/function"
)

const payloadFile = "../../payload.json"

func main() {
	ctx := context.Background()

	payload := startPayloadFromFile()

	HandleRequest, err := handler.HandleRequest(ctx, payload)

	if err != nil {
		log.Fatalln("Handle request failure:\n", err)
	}

	log.Printf("%+v", HandleRequest)
}

func startPayloadFromFile() f.EventPayload {
	file, err := ioutil.ReadFile(payloadFile)

	if err != nil {
		log.Fatalln("Payload read file error:\n", err)
	}

	payload := f.EventPayload{}

	err = json.Unmarshal([]byte(file), &payload)

	if err != nil {
		log.Fatalln("Failed to unmarshal payload:\n", err)
	}

	return payload
}
