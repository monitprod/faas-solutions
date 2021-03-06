package local

import (
	"context"
	"encoding/json"
	"io/ioutil"

	log "github.com/sirupsen/logrus"

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

// TODO: Document this
func Start(
	b StartBuilder,
) {
	// TODO: Document this
	ctx := context.WithValue(context.Background(), c.IsLocal, true)
	ctx = context.WithValue(ctx, c.LocalMainFunc, b.LocalFunc)

	if b.Payload == nil {
		b.Payload = startPayloadFromFile(b.PayloadFile)
	}

	HandleRequest, err := b.Handler(ctx, *b.Payload)

	if err != nil {
		log.Errorln("Handle request failure:\n", err)
	}

	log.Printf("%+v", HandleRequest)
}

func startPayloadFromFile(payloadFile string) *map[string]interface{} {

	file, err := ioutil.ReadFile(payloadFile)

	if err != nil {
		log.Errorln("Payload read file error:\n", err)
	}

	payload := make(map[string]interface{})

	err = json.Unmarshal([]byte(file), &payload)

	if err != nil {
		log.Errorln("Failed to unmarshal payload:\n", err)
	}

	return &payload
}
