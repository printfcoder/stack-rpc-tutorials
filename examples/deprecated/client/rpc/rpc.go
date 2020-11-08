package main

import (
	"context"
	"fmt"

	proto "github.com/micro-in-cn/tutorials/examples/client/rpc/proto"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/errors"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/transport/grpc"
)

type Example struct{}

type Foo struct{}

func (e *Example) Call(ctx context.Context, req *proto.CallRequest, rsp *proto.CallResponse) error {
	logger.Infof("收到 Example.Call 请求 %v\n", req)
	fmt.Printf("%v\n", req)

	if len(req.Name) == 0 {
		return errors.BadRequest("go.micro.rpc.example", "no content")
	}

	rsp.Message = "RPC Call收到了你的请求 " + req.Name
	return nil
}

func (f *Foo) Bar(ctx context.Context, req *proto.EmptyRequest, rsp *proto.EmptyResponse) error {
	logger.Infof("收到 Foo.Bar 请求")
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.rpc.example"),
		micro.Transport(grpc.NewTransport()),
	)

	service.Init()

	// 注册 example 接口
	proto.RegisterExampleHandler(service.Server(), new(Example))

	// 注册 foo 接口
	proto.RegisterFooHandler(service.Server(), new(Foo))

	if err := service.Run(); err != nil {
		logger.Error(err)
	}
}
