package main

import (
	"github.com/influxdata/influxdb/services/subscriber"
	"github.com/micro-in-cn/micro-tutorials/microservice-in-micro/part1/user-service/handler"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"

	s "github.com/micro-in-cn/micro-tutorials/microservice-in-micro/part1/user-service/proto/service"
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
	s.RegisterServiceHandler(service.Server(), new(handler.Service))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("mu.micro.book.user.srv.service", service.Server(), new(subscriber.Service))

	// Register Function as Subscriber
	micro.RegisterSubscriber("mu.micro.book.user.srv.service", service.Server(), subscriber.Service)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
