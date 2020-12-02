package main

import (
	"context"
	"fmt"

	"github.com/stack-labs/stack-rpc"
	proto "github.com/stack-labs/stack-rpc-tutorials/examples/proto/service/rpc"
	"github.com/stack-labs/stack-rpc/server"
	"github.com/stack-labs/stack-rpc/util/log"
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
	service := stack.NewService(
		stack.Name("greeter"),
		// 声明包装器
		stack.WrapHandler(logWrapper1, logWrapper2),
	)

	service.Init()

	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
