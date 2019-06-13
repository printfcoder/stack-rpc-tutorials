package main

import (
	"fmt"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part4/basic"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part4/basic/common"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part4/basic/config"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part4/user-srv/handler"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part4/user-srv/model"
	"github.com/micro/cli"
	"github.com/micro/go-config/source/grpc"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"time"

	s "github.com/micro-in-cn/tutorials/microservice-in-micro/part4/user-srv/proto/user"
)

var (
	appName = "user_srv"
	cfg     = &userCfg{}
)

type userCfg struct {
	common.AppCfg
}

func main() {
	// 初始化配置、数据库等信息
	initCfg()

	// 使用consul注册
	micReg := consul.NewRegistry(registryOptions)

	// 新建服务
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
	s.RegisterUserHandler(service.Server(), new(handler.Service))

	// 启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	consulCfg := &common.Consul{}
	err := config.C().App("consul", consulCfg)
	if err != nil {
		panic(err)
	}

	ops.Timeout = time.Second * 5
	ops.Addrs = []string{fmt.Sprintf("%s:%d", consulCfg.Host, consulCfg.Port)}
}

func initCfg() {
	source := grpc.NewSource(
		grpc.WithAddress("127.0.0.1:9600"),
		grpc.WithPath("micro"),
	)

	basic.Init(config.WithSource(source))

	err := config.C().App(appName, cfg)
	if err != nil {
		panic(err)
	}

	log.Logf("[initCfg] 配置，cfg：%v", cfg)

	return
}
