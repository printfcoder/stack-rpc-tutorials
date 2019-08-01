package subscriber

import (
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/util/log"
)

func Handler(event broker.Event) error {
	log.Logf("[sub] 收到消息，请查收: %v, %v", event.Message().Header, string(event.Message().Body))
	return nil
}
