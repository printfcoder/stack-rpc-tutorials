package subscriber

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/broker"
	_ "github.com/micro/go-plugins/broker/nats"
)

func Handler(event broker.Event) error {
	log.Logf("[sub] received messages: %v, %v", event.Message().Header, string(event.Message().Body))
	return nil
}
