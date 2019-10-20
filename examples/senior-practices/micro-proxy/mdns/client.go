package main

import (
	"context"
	"fmt"

	proto "github.com/micro-in-cn/tutorials/examples/basic-practices/micro-service/proto"
	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(micro.Name("proxy.consul.client"))
	service.Init()

	// 创建客户端
	greeter := proto.NewGreeterService("proxy.consul.service", service.Client())

	// 调用greeter服务
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "Micro中国"})
	if err != nil {
		fmt.Println(err)
		return
	}

	// 打印响应结果
	fmt.Println(rsp.Greeting)

	// 不需要启动服务，直接退出
	// if err := service.Run(); err != nil {
	// 	fmt.Println(err)
	// }
}
