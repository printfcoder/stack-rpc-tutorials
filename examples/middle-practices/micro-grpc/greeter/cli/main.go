package main

import (
	"context"
	"fmt"

	pb "github.com/micro-in-cn/tutorials/examples/middle-practices/micro-grpc/proto/go/micro"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/service/grpc"
)

func main() {
	service := grpc.NewService()
	service.Init()

	// use the generated client stub
	cl := pb.NewSayService("go.micro.srv.greeter", service.Client())

	// Set arbitrary headers in context
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id": "John",
		"X-From-Id": "script",
	})

	rsp, err := cl.Hello(ctx, &pb.Request{
		Name: "我是来自micro风格的客户端请求",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Msg)
}
