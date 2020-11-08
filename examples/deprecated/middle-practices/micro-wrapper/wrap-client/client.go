package main

import (
	"context"

	proto "github.com/micro-in-cn/tutorials/examples/basic-practices/micro-service/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/util/log"
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
	service := micro.NewService(
		micro.Name("wrap.client.cli"),
		micro.WrapClient(
			NewClientWrapper(),
		),
	)

	cl := proto.NewGreeterService("wrap.call.service", service.Client())
	rsp, err := cl.Hello(context.Background(), &proto.HelloRequest{Name: "Micro中国"})
	if err != nil {
		panic(err)
	}

	log.Info(rsp.Greeting)
}
