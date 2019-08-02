package main

import (
	"time"

	"github.com/micro-in-cn/tutorials/micro-sync/inventory-srv/handler"
	"github.com/micro-in-cn/tutorials/micro-sync/inventory-srv/model"
	proto "github.com/micro-in-cn/tutorials/micro-sync/inventory-srv/proto/inventory"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
)

var (
	appName = "inv_srv"
)

func main() {
	// 新建服务
	service := micro.NewService(
		micro.Name("go.micro.srv.sync.inv"),
		micro.RegisterTTL(time.Second*15),
		micro.RegisterInterval(time.Second*10),
	)

	// 服务初始化
	service.Init(
		micro.Action(func(c *cli.Context) {
			// 初始化模型层
			model.Init()
			// 初始化handler
			handler.Init()
		}),
	)

	// 注册服务
	proto.RegisterInventoryHandler(service.Server(), new(handler.Service))

	// 启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
