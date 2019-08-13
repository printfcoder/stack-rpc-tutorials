package main

import (
	"flag"

	"github.com/micro-in-cn/tutorials/micro-benchmark/micro/internal"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	reg "github.com/micro/go-plugins/registry/nats"
	"github.com/micro/go-plugins/transport/nats"
	nats2 "github.com/nats-io/nats.go"
)

var concurrency = flag.Int("c", 1, "concurrency")
var total = flag.Int("n", 1, "total requests for all clients")

func main() {
	flag.Parse()
	n := *concurrency
	m := *total / n

	opts := nats2.GetDefaultOptions()
	// 替换为具体的地址
	opts.Servers = []string{"127.0.0.1:4222"}
	r := reg.NewRegistry(func(ops *registry.Options) {
		ops.Addrs = []string{"127.0.0.1:4222"}
	})

	internal.ClientRun(m, n,
		"go.micro.benchmark.hello.nats_transport",
		func() client.Client {
			return client.NewClient(
				client.Transport(nats.NewTransport(nats.Options(opts))),
				client.Registry(r),
			)
		})
}
