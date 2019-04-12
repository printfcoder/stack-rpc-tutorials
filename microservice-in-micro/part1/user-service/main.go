package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro-in-cn/micro-tutorials/microservice-in-micro/part1/user-service/handler"
	"github.com/micro-in-cn/micro-tutorials/microservice-in-micro/part1/user-service/subscriber"

	example "github.com/micro-in-cn/micro-tutorials/microservice-in-micro/part1/user-service/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("mu.micro.book.user.srv.service"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("mu.micro.book.user.srv.service", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("mu.micro.book.user.srv.service", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
