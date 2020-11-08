package main

import (
	"context"
	"fmt"

	proto "github.com/micro-in-cn/tutorials/examples/service/proto"
	"github.com/micro/go-micro/v2"
)

func main() {
	// 定义服务，可以传入其它可选参数
	service := micro.NewService(micro.Name("greeter.client"))
	service.Init()

	// 创建客户端
	greeter := proto.NewGreeterService("greeter.service", service.Client())

	// 调用greeter服务
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "Micro中国"})
	if err != nil {
		fmt.Println(err)
		return
	}

	// 打印响应结果
	fmt.Println(rsp.Greeting)
}
