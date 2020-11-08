package main

import (
	"context"
	"fmt"
	"time"

	proto "github.com/micro-in-cn/tutorials/examples/basic-practices/micro-service/proto"
	"github.com/micro/go-micro/v2"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	time.Sleep(time.Second * 9)
	rsp.Greeting = "你好，" + req.Name
	return nil
}

func main() {

	// 创建服务，除了服务名，其它选项可加可不加，比如Version版本号、Metadata元数据等
	service := micro.NewService(
		micro.Name("timeout.service"),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type": "你好，世界",
		}),
	)
	service.Init()

	// 注册服务
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	// 启动服务
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
