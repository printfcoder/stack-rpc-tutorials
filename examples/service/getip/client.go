package main

import (
	"context"

	"github.com/stack-labs/stack-rpc"
	proto "github.com/stack-labs/stack-rpc-tutorials/examples/proto/service/rpc"
	"github.com/stack-labs/stack-rpc/logger"
)

func main() {
	// 定义服务，可以传入其它可选参数
	service := stack.NewService(stack.Name("stack.rpc.client"))
	service.Init()

	// 创建客户端
	greeter := proto.NewGreeterService("stack.rpc.greeter", service.Client())

	// 调用greeter服务
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "StackLabs"})
	if err != nil {
		logger.Fatal(err)
		return
	}

	// 打印响应结果
	logger.Info(rsp.Greeting)
}
