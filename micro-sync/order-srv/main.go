package main

import (
	"time"

	"github.com/micro-in-cn/tutorials/micro-sync/order-srv/handler"
	"github.com/micro-in-cn/tutorials/micro-sync/order-srv/model"
	proto "github.com/micro-in-cn/tutorials/micro-sync/order-srv/proto/orders"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
)

func main() {
	// 新建服务
	service := micro.NewService(
		micro.Name("go.micro.srv.sync.order"),
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
	err := proto.RegisterOrdersHandler(service.Server(), new(handler.Orders))
	if err != nil {
		log.Fatal(err)
	}

	// 启动服务
	if err = service.Run(); err != nil {
		log.Fatal(err)
	}
}
