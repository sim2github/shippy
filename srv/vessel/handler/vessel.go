package handler

import (
	"context"

	"github.com/micro/go-micro/v2/util/log"

	vessel "github.com/EwanValentine/shippy/srv/vessel/proto/vessel"
)

type Vessel struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Vessel) Call(ctx context.Context, req *vessel.Request, rsp *vessel.Response) error {
	log.Log("Received Vessel.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Vessel) Stream(ctx context.Context, req *vessel.StreamingRequest, stream vessel.Vessel_StreamStream) error {
	log.Logf("Received Vessel.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&vessel.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Vessel) PingPong(ctx context.Context, stream vessel.Vessel_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&vessel.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
