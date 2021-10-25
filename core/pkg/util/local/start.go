package local

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	c "github.com/monitprod/core/pkg/constant"
)

type LocalFunc func(payload *map[string]interface{})

type HandleRequest func(
	ctx context.Context,
	payload map[string]interface{},
) (
	response map[string]interface{},
	err error,
)

type StartBuilder struct {
	Handler     HandleRequest
	LocalFunc   LocalFunc
	Payload     *map[string]interface{}
	PayloadFile string // If Payload exists, it will not be read
}

func Start(
	b StartBuilder,
) {
	ctx := context.WithValue(context.Background(), c.IsLocal, true)
	ctx = context.WithValue(ctx, c.LocalMainFunc, b.LocalFunc)

	if b.Payload == nil {
		b.Payload = startPayloadFromFile(b.PayloadFile)
	}

	HandleRequest, err := b.Handler(ctx, *b.Payload)

	if err != nil {
		log.Fatalln("Handle request failure:\n", err)
	}

	log.Printf("%+v", HandleRequest)
}

func startPayloadFromFile(payloadFile string) *map[string]interface{} {

	file, err := ioutil.ReadFile(payloadFile)

	if err != nil {
		log.Fatalln("Payload read file error:\n", err)
	}

	payload := make(map[string]interface{})

	err = json.Unmarshal([]byte(file), &payload)

	if err != nil {
		log.Fatalln("Failed to unmarshal payload:\n", err)
	}

	return &payload
}
