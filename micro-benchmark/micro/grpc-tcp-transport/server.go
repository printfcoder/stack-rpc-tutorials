package main

import (
	"fmt"
	"github.com/micro/go-micro/service/grpc"

	"github.com/micro-in-cn/tutorials/micro-benchmark/micro/internal"
	"github.com/micro-in-cn/tutorials/micro-benchmark/pb"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/transport/tcp"
)

func main() {
	service := grpc.NewService(
		micro.Name("go.micro.benchmark.hello.grpc_tcp"),
		micro.Version("latest"),
		micro.Transport(tcp.NewTransport()),
	)

	service.Init()

	pb.RegisterHelloHandler(service.Server(), &internal.HelloS{})

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
