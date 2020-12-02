package main

import (
	"context"
	"fmt"

	"github.com/stack-labs/stack-rpc"
	proto "github.com/stack-labs/stack-rpc-tutorials/examples/proto/service/rpc"
	"github.com/stack-labs/stack-rpc/client"
)

// logWrapper1 打印每次请求的信息
type logWrapper1 struct {
	client.Client
}

// logWrapper2 打印每次请求的信息
type logWrapper2 struct {
	client.Client
}

func (l *logWrapper1) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	fmt.Printf("[Call1] 客户端请求服务：%s，方法：%s\n", req.Service(), req.Endpoint())
	return l.Client.Call(ctx, req, rsp)
}

func (l *logWrapper2) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	fmt.Printf("[Call2] 客户端请求服务：%s，方法：%s\n", req.Service(), req.Endpoint())
	return l.Client.Call(ctx, req, rsp)
}

func logWrap1(c client.Client) client.Client {
	return &logWrapper1{c}
}

func logWrap2(c client.Client) client.Client {
	return &logWrapper2{c}
}

func main() {
	service := stack.NewService(
		stack.Name("greeter.client"),
		// 把客户端包装起来，包装器执行顺序与声明顺序有关
		stack.WrapClient(logWrap2, logWrap1),
	)

	service.Init()

	greeter := proto.NewGreeterService("greeter", service.Client())

	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "StackLabs"})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Greeting)
}
