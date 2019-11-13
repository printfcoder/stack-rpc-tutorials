package main

import (
	"context"
	"fmt"
	"time"

	proto "github.com/micro-in-cn/tutorials/examples/basic-practices/micro-service/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
)

func main() {
	cli := client.NewClient(
		client.RequestTimeout(time.Second * 15),
	)

	// 定义服务，可以传入其它可选参数
	service := micro.NewService(
		micro.Name("timeout.client"),
		micro.Client(cli))

	service.Init()

	// 创建客户端
	greeter := proto.NewGreeterService("timeout.service", service.Client())

	// 调用greeter服务
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "Micro中国"})
	if err != nil {
		fmt.Println(err)
		return
	}

	// 打印响应结果
	fmt.Println(rsp.Greeting)
}
