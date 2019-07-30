package main

import (
	"context"
	"fmt"

	proto "github.com/micro-in-cn/tutorials/examples/basic-practices/micro-service/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "你好，" + req.Name
	return nil
}

func main() {
	// 新建broker
	bk1 := broker.NewBroker(
		broker.Addrs(fmt.Sprintf("%s:%d", "127.0.0.1", 11089)),
	)

	// 订阅主题1
	_, err := bk1.Subscribe("go.micro.topic.custom1", func(p broker.Publication) error {
		fmt.Println("[bk1] 订阅收到主题1消息：", string(p.Message().Body), "header", p.Message().Header)
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// 订阅主题2
	_, err = bk1.Subscribe("go.micro.topic.custom2", func(p broker.Publication) error {
		fmt.Println("[bk1] 订阅收到主题2消息：", string(p.Message().Body), "header", p.Message().Header)
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// 创建服务
	service := micro.NewService(
		micro.Name("broker.service"),
		micro.Broker(
			bk1,
		),
	)

	// 注册服务
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	// 启动服务
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
