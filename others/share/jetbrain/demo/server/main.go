package main

import (
	"context"

	goland "github.com/micro-in-cn/tutorials/others/share/jetbrain/demo/proto"
	"github.com/micro/go-micro/v2"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *goland.HelloRequest, rsp *goland.HelloResponse) error {
	rsp.Greeting = "Hello，" + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("goland.greeter"),
	)
	service.Init()
	goland.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		panic(err)
	}
}
