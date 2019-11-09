package main

import (
	"fmt"

	"github.com/micro-in-cn/tutorials/micro-benchmark/micro/internal"
	"github.com/micro-in-cn/tutorials/micro-benchmark/pb"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/transport"
	"github.com/micro/go-plugins/registry/consul"
	"github.com/micro/go-plugins/transport/rabbitmq"
)

func main() {
	// 使用consul注册
	micReg := consul.NewRegistry(func(ops *registry.Options) {
		ops.Addrs = []string{"127.0.0.1:8500"}
	})

	service := micro.NewService(
		micro.Name("go.micro.benchmark.hello.rabbitmq_transport"),
		micro.Version("latest"),
		micro.Registry(micReg),
		micro.Transport(rabbitmq.NewTransport(transport.Addrs("amqp://guest:guest@127.0.0.1:5672"))),
	)

	service.Init()

	pb.RegisterHelloHandler(service.Server(), &internal.HelloS{})

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
