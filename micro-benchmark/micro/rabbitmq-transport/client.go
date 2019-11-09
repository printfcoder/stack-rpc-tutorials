package main

import (
	"flag"
	"github.com/micro-in-cn/tutorials/micro-benchmark/micro/internal"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/transport"
	"github.com/micro/go-plugins/registry/consul"
	"github.com/micro/go-plugins/transport/rabbitmq"
)

var concurrency = flag.Int("c", 1, "concurrency")
var total = flag.Int("n", 1, "total requests for all clients")

func main() {
	flag.Parse()
	n := *concurrency
	m := *total / n

	// 使用consul注册
	micReg := consul.NewRegistry(func(ops *registry.Options) {
		ops.Addrs = []string{"127.0.0.1:8500"}
	})

	internal.ClientRun(m, n, "go.micro.benchmark.hello.rabbitmq_transport", func() client.Client {
		return client.NewClient(
			client.Registry(micReg),
			client.Transport(rabbitmq.NewTransport(transport.Addrs("amqp://guest:guest@127.0.0.1:5672"))),
		)
	})
}
