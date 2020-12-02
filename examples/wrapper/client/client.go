package main

import (
	"context"

	"github.com/stack-labs/stack-rpc"
	proto "github.com/stack-labs/stack-rpc-tutorials/examples/proto/service/rpc"
	"github.com/stack-labs/stack-rpc/client"
	log "github.com/stack-labs/stack-rpc/logger"
)

type logWrapper struct {
	client.Client
}

func (l *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	log.Infof("[Call] 请求服务：%s.%s", req.Service(), req.Endpoint())
	return l.Client.Call(ctx, req, rsp)
}

func NewClientWrapper() client.Wrapper {
	return func(cli client.Client) client.Client {
		return &logWrapper{cli}
	}
}

func main() {
	service := stack.NewService(
		stack.Name("wrap.client.cli"),
		stack.WrapClient(
			NewClientWrapper(),
		),
	)

	cl := proto.NewGreeterService("wrap.call.service", service.Client())
	rsp, err := cl.Hello(context.Background(), &proto.HelloRequest{Name: "StackLabs"})
	if err != nil {
		panic(err)
	}

	log.Info(rsp.Greeting)
}
