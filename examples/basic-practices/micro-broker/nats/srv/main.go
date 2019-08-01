package main

import (
	"github.com/brucewangno1/go-micro-pubsub-with-nats/srv/subscriber"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	_ "github.com/micro/go-plugins/broker/nats"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.pubsub"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	broker := service.Server().Options().Broker

	if err := broker.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := broker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}

	if _, err := broker.Subscribe("go.micro.pubsub.topic.event", subscriber.Handler); err != nil {
		log.Fatalf("broker.Subscribe error: %v", err)
	}

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
