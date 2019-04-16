package main

import (
	"fmt"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part2/basic"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part2/basic/config"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part2/user-service/handler"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part2/user-service/model"
	"github.com/micro/cli"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"time"

	s "github.com/micro-in-cn/tutorials/microservice-in-micro/part2/user-service/proto/service"
)

func main() {

	// 初始化配置、数据库等信息
	basic.Init()

	// 使用consul注册
	micReg := consul.NewRegistry(registryOptions)

	// New Service
	service := micro.NewService(
		micro.Name("mu.micro.book.srv.user"),
		micro.Registry(micReg),
		micro.Version("latest"),
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
	s.RegisterServiceHandler(service.Server(), new(handler.Service))

	// 启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	consulCfg := config.GetConsulConfig()
	ops.Timeout = time.Second * 5
	ops.Addrs = []string{fmt.Sprintf("%s:%d", consulCfg.GetHost(), consulCfg.GetPort())}
}
