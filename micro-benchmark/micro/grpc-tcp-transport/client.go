package main

import (
	"flag"
	"github.com/micro-in-cn/tutorials/micro-benchmark/micro/internal"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/grpc"
	"github.com/micro/go-plugins/transport/tcp"
)

var concurrency = flag.Int("c", 1, "concurrency")
var total = flag.Int("n", 1, "total requests for all clients")

func main() {
	flag.Parse()
	n := *concurrency
	m := *total / n

	internal.ClientRun(m, n, "go.micro.benchmark.hello.grpc_tcp",
		func() client.Client {
			return grpc.NewClient(
				client.Transport(tcp.NewTransport()),
			)
		},
	)
}
