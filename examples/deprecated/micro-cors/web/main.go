package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/transport/grpc"
	"github.com/micro/go-micro/v2/web"
	"log"
)

func main() {
	// New web service
	service := web.NewService(
		web.Name("go.micro.web.cors"),
		web.Address(":9090"),
		web.MicroService(micro.NewService(micro.Transport(grpc.NewTransport()))),
	)

	service.Options().Service.Client()

	if err := service.Init(); err != nil {
		log.Fatal("Init", err)
	}

	if err := service.Run(); err != nil {
		log.Fatal("Run: ", err)
	}
}
