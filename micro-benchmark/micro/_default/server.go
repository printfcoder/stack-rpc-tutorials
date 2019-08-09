package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"

	"github.com/micro-in-cn/tutorials/micro-benchmark/pb"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
)

type HelloS struct{}

func (t *HelloS) Say(ctx context.Context, args *pb.BenchmarkMessage, reply *pb.BenchmarkMessage) error {
	s := "OK"
	var i int32 = 100
	args.Field1 = s
	args.Field2 = i
	*reply = *args
	if *delay > 0 {
		time.Sleep(*delay)
	} else {
		runtime.Gosched()
	}
	return nil
}

var delay = flag.Duration("delay", 0, "delay to mock business processing")

func main() {
	//flag.Parse()

	service := micro.NewService(
		micro.Name("go.micro.benchmark.hello"),
		micro.Version("latest"),
	)

	service.Init()

	pb.RegisterHelloHandler(service.Server(), &HelloS{})

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
