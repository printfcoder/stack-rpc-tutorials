package main

import (
	"flag"
	"github.com/micro-in-cn/tutorials/micro-benchmark/micro/internal"
	"github.com/micro-in-cn/tutorials/micro-benchmark/pb"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/transport/tcp"
)

var concurrency = flag.Int("c", 1, "concurrency")
var total = flag.Int("n", 1, "total requests for all clients")

func main() {
	flag.Parse()
	n := *concurrency
	m := *total / n

	service := micro.NewService(micro.Name("go.micro.benchmark.hello.client_tcp"), micro.Transport(tcp.NewTransport()), )
	c := pb.NewHelloService("go.micro.benchmark.hello.default_tcp", service.Client(), )

	internal.ClientRun(m, n, c)
}
