package publisher

import (
	"context"
	"github.com/micro/go-log"
)

func Handler(ctx context.Context, msg *example.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
