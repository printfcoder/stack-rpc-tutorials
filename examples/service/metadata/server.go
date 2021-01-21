package main

import (
	"context"
	"fmt"

	"github.com/stack-labs/stack-rpc"
	"github.com/stack-labs/stack-rpc-plugins/logger/logrus"
	proto "github.com/stack-labs/stack-rpc-tutorials/examples/proto/service/rpc"
	"github.com/stack-labs/stack-rpc/logger"
	"github.com/stack-labs/stack-rpc/pkg/metadata"
)

// 服务类
type Greeter struct{}

// 实现proto中的Hello接口
func (g Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		rsp.Greeting = fmt.Sprintf("Hi! %s No metadata received.", req.Name)
		return nil
	}

	rsp.Greeting = fmt.Sprintf("Hello %s thanks for this %v, %v", req.Name, md["User"], md["Id"])
	return nil
}

func main() {
	// 实例化服务，并命名为stack.rpc.greeter
	service := stack.NewService(
		stack.Name("stack.rpc.greeter"),
		stack.Logger(logrus.NewLogger()),
	)
	// 初始化服务
	service.Init()

	// 将Greeter注册到服务上
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	// 运行服务
	if err := service.Run(); err != nil {
		logger.Error(err)
	}
}
