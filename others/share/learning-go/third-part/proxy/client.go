package main

import (
	"context"

	proto "github.com/micro-in-cn/tutorials/others/share/learning-go/third-part/proto/learning"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
)

func main() {
	service := micro.NewService(micro.Name("go.micro.client.greeter"))
	service.Init()

	// 创建客户端
	greeter := proto.NewGreeterService("go.micro.api.v2.learning", service.Client())

	// 调用greeter服务
	rsp, err := greeter.Hi(context.TODO(), &proto.Request{Name: "Micro"})
	if err != nil {
		log.Error(err)
		return
	}

	// 打印响应结果
	log.Info(rsp.Msg)

	// 不用运行
	// service.Run()
}
