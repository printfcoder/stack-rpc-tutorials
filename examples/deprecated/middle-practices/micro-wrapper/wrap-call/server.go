package main

import (
	"context"
	"fmt"

	proto "github.com/micro-in-cn/tutorials/examples/basic-practices/micro-service/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
	"github.com/prometheus/common/log"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	newMd, _ := metadata.FromContext(ctx)
	rsp.Greeting = "Hi! " + req.Name

	log.Info("[Hello] call-wrapped1: ", newMd["Call-Wrapped1"])
	log.Info("[Hello] call-wrapped2: ", newMd["Call-Wrapped2"])
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("wrap.call.service"),
	)
	service.Init()

	// 注册服务
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	// 启动服务
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
