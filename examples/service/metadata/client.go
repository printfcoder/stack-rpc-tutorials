package main

import (
	"context"

	"github.com/stack-labs/stack-rpc"
	proto "github.com/stack-labs/stack-rpc-tutorials/examples/proto/service/rpc"
	"github.com/stack-labs/stack-rpc/logger"
	"github.com/stack-labs/stack-rpc/pkg/metadata"
)

func main() {
	// 定义服务，可以传入其它可选参数
	service := stack.NewService(stack.Name("stack.rpc.client"))
	service.Init()

	// 创建客户端
	greeter := proto.NewGreeterService("stack.rpc.greeter", service.Client())

	// Set 设置上下文
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"User": "StackDemo",
		// 注意，Stack会强制将metadata中的字段大驼峰化，也即ID会为转成Id
		"ID": "1",
	})

	// 调用greeter服务
	rsp, err := greeter.Hello(ctx, &proto.HelloRequest{Name: "StackLabs"})
	if err != nil {
		logger.Fatal(err)
		return
	}

	// 打印响应结果
	logger.Info(rsp.Greeting)
}
