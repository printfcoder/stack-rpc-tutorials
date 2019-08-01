package main

import (
	"github.com/brucewangno1/go-micro-pubsub-with-nats/cli/publisher"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.cli.pubsub"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	b := service.Client().Options().Broker

	if err := b.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := b.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}

	go publisher.SendEvent("go.micro.pubsub.topic.event", b)
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
