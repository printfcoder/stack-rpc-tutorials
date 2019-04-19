package subscriber

import (
	"context"
	"github.com/micro/go-log"

	example "github.com/micro-in-cn/tutorials/microservice-in-micro/part3/payment-srv/proto/example"
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
