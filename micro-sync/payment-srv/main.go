package main

import (
	"time"

	"github.com/micro-in-cn/tutorials/micro-sync/payment-srv/handler"
	"github.com/micro-in-cn/tutorials/micro-sync/payment-srv/model"
	s "github.com/micro-in-cn/tutorials/micro-sync/payment-srv/proto/payment"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
)

func main() {
	// 新建服务
	service := micro.NewService(
		micro.Name("go.micro.sync.srv.payment"),
		micro.RegisterTTL(time.Second*15),
		micro.RegisterInterval(time.Second*10),
	)

	// 服务初始化
	service.Init(
		micro.Action(func(c *cli.Context) {
			// 初始化模型层
			model.Init(service.Options())
			// 初始化handler
			handler.Init()
		}),
	)

	// 注册服务
	s.RegisterPaymentHandler(service.Server(), new(handler.Service))

	// 启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
