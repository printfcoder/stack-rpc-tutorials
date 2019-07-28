package main

import (
	"context"

	"time"

	proto "github.com/micro-in-cn/tutorials/examples/senior-practices/micro-proxy/grpc/api/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/util/log"
)

type Greeter struct{}

// Example.SayHello 通过API向外暴露为/example/sayhello，接收http请求
// 即：/example/Say请求会调用go.micro.api.example服务的Example.Call方法
func (e *Greeter) SayHello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloReply) error {
	log.Log("Greeter.Call接口收到请求")
	rsp.Message = "Hello! " + req.Name
	return nil
}

func main() {
	service := grpc.NewService(
		micro.Name("greeter"),
		micro.Address("127.0.0.1:10000"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	service.Init()

	// 注册 example handler
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
