package main

import (
	"context"
	"log"

	proto "github.com/micro-in-cn/tutorials/examples/basic-practices/micro-api/rpc/proto"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/errors"
)

type Example struct{}

type Foo struct{}

// Call 方法会接收由API层转发，路由为/example/call的HTTP请求
func (e *Example) Call(ctx context.Context, req *proto.CallRequest, rsp *proto.CallResponse) error {
	log.Print("收到 Example.Call 请求")

	if len(req.Name) == 0 {
		return errors.BadRequest("go.micro.api.example", "no content")
	}

	rsp.Message = "RPC Call收到了你的请求 " + req.Name
	return nil
}

// Bar 方法会接收由API层转发，路由为/example/foo/bar的HTTP请求
// 该接口我们什么参数也不处理，只打印信息
func (f *Foo) Bar(ctx context.Context, req *proto.EmptyRequest, rsp *proto.EmptyResponse) error {
	log.Print("收到 Foo.Bar 请求")
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.example"),
	)

	service.Init()

	// 注册 example 接口
	proto.RegisterExampleHandler(service.Server(), new(Example))

	// 注册 foo 接口
	proto.RegisterFooHandler(service.Server(), new(Foo))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
