package main

import (
	"fmt"

	"context"

	proto "github.com/micro-in-cn/tutorials/examples/basic-practices/micro-service/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/util/log"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "你好呀！" + req.Name
	return nil
}

// logWrapper1 包装HandlerFunc类型的接口
func logWrapper1(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		log.Logf("[logWrapper1] %s 收到请求", req.Endpoint())
		err := fn(ctx, req, rsp)
		return err
	}
}

// logWrapper2 包装HandlerFunc类型的接口
func logWrapper2(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		log.Logf("[logWrapper2] %s 收到请求", req.Endpoint())
		err := fn(ctx, req, rsp)
		return err
	}
}

func main() {
	service := micro.NewService(
		micro.Name("greeter"),
		// 声明包装器
		micro.WrapHandler(logWrapper1, logWrapper2),
	)

	service.Init()

	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
