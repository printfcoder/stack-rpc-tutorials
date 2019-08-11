package main

import (
	"fmt"
	"time"

	"github.com/micro-in-cn/tutorials/micro-benchmark/micro/internal"
	"github.com/micro-in-cn/tutorials/micro-benchmark/pb"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
)

func main() {
	service := grpc.NewService(
		micro.Name("go.micro.benchmark.hello.grpc"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	service.Init()

	pb.RegisterHelloHandler(service.Server(), &internal.HelloS{})

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
