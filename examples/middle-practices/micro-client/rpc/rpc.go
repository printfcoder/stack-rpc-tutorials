package main

import (
	"context"
	"fmt"
	"log"

	proto "github.com/micro-in-cn/tutorials/examples/middle-practices/micro-client/rpc/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
)

type Example struct{}

type Foo struct{}

func (e *Example) Call(ctx context.Context, req *proto.CallRequest, rsp *proto.CallResponse) error {
	log.Printf("收到 Example.Call 请求 %v\n", req)
	fmt.Printf("%v\n", req)

	if len(req.Name) == 0 {
		return errors.BadRequest("go.micro.rpc.example", "no content")
	}

	rsp.Message = "RPC Call收到了你的请求 " + req.Name
	return nil
}

func (f *Foo) Bar(ctx context.Context, req *proto.EmptyRequest, rsp *proto.EmptyResponse) error {
	log.Print("收到 Foo.Bar 请求")
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.rpc.example"),
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
