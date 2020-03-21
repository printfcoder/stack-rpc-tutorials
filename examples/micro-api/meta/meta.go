package main

import (
	"context"
	"log"

	proto "github.com/micro-in-cn/tutorials/examples/micro-api/rpc/proto"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/api"
	rapi "github.com/micro/go-micro/v2/api/handler/api"
	"github.com/micro/go-micro/v2/api/handler/rpc"
	"github.com/micro/go-micro/v2/errors"
)

type Example struct{}

type Foo struct{}

// Call 方法在下面main中我们通过endpoint将其注册到/example/call路由
func (e *Example) Call(ctx context.Context, req *proto.CallRequest, rsp *proto.CallResponse) error {
	log.Print("Meta Example.Call接口收到请求")

	if len(req.Name) == 0 {
		return errors.BadRequest("go.micro.api.example", "no content")
	}

	rsp.Message = "Meta已经收到你的请求，" + req.Name
	return nil
}

// Bar 方法在下面main中我们通过endpoint将其注册到/example/foo/bar路由
func (f *Foo) Bar(ctx context.Context, req *proto.EmptyRequest, rsp *proto.EmptyResponse) error {
	log.Print("Meta Foo.Bar接口收到请求")
	// noop

	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.example"),
	)

	service.Init()

	// 注册Example接口处理器
	proto.RegisterExampleHandler(service.Server(), new(Example), api.WithEndpoint(&api.Endpoint{
		// 接口方法，一定要在proto接口中存在，不能是类的自有方法
		Name: "Example.Call",
		// http请求路由，支持POSIX风格
		Path: []string{"/example"},
		// 支持的方法类型
		Method: []string{"POST", "GET"},
		// 该接口使用的API转发模式
		Handler: rpc.Handler,
	}))

	// 注册Foo接口处理器
	proto.RegisterFooHandler(service.Server(), new(Foo), api.WithEndpoint(&api.Endpoint{
		Name:    "Foo.Bar",
		Path:    []string{"/foo/bar"},
		Method:  []string{"POST", "GET", "DELETE"},
		Handler: rapi.Handler,
	}))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
