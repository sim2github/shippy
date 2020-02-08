package subscriber

import (
	"context"
	"github.com/micro/go-micro/v2/util/log"

	vessel "github.com/EwanValentine/shippy/srv/vessel/proto/vessel"
)

type Vessel struct{}

func (e *Vessel) Handle(ctx context.Context, msg *vessel.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *vessel.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
