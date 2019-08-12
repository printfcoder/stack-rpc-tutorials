package main

import (
	"fmt"

	"github.com/micro-in-cn/tutorials/micro-benchmark/micro/internal"
	"github.com/micro-in-cn/tutorials/micro-benchmark/pb"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	reg "github.com/micro/go-plugins/registry/nats"
	"github.com/micro/go-plugins/transport/nats"
	nats2 "github.com/nats-io/nats.go"
)

func main() {
	opts := nats2.GetDefaultOptions()
	// 替换为具体的地址
	opts.Servers = []string{"127.0.0.1:4222"}
	t := nats.NewTransport(nats.Options(opts))

	r := reg.NewRegistry(func(ops *registry.Options) {
		ops.Addrs = []string{"127.0.0.1:4222"}
	})

	service := micro.NewService(
		micro.Name("go.micro.benchmark.hello.nats_transport"),
		micro.Version("latest"),
		micro.Transport(t),
		micro.Registry(r),
	)

	service.Init()

	pb.RegisterHelloHandler(service.Server(), &internal.HelloS{})

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
