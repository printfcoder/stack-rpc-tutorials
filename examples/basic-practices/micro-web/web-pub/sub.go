package main

import (
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/util/log"
)

var (
	topic = "go.micro.web.topic.hi"
)

func main() {

	bk := broker.NewBroker(
		broker.Addrs(fmt.Sprintf("%s:%d", "127.0.0.1", 11089)),
	)

	_, err := bk.Subscribe(topic, func(p broker.Publication) error {
		log.Logf("[sub] Received Body: %s, Header: %s", string(p.Message().Body), p.Message().Header)
		return nil
	})
	if err != nil {
		log.Logf("[ERR] err: %s", err)
	}

	s := micro.NewService(
		micro.Name("go.micro.book.web.sub"),
		micro.Version("latest"),
		micro.Address(":8089"),
		micro.Broker(bk),
	)

	s.Init()

	_ = s.Run()
}
