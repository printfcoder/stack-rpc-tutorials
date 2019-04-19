package main

import (
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part3/payment-srv/handler"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part3/payment-srv/subscriber"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"

	example "github.com/micro-in-cn/tutorials/microservice-in-micro/part3/payment-srv/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("mu.micro.book.srv.payment"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("mu.micro.book.srv.payment", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("mu.micro.book.srv.payment", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
