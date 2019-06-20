package subscriber

import (
	"github.com/micro/go-log"

	"context"
	example "github.com/micro-in-cn/tutorials/micro-istio/http/srv/proto/example"
)

type Example struct{}

func (e *Example) Handle(ctx context.Context, msg *example.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *example.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
