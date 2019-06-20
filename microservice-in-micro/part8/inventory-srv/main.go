package main

import (
	"fmt"
	"time"

	"github.com/micro-in-cn/tutorials/microservice-in-micro/part8/basic"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part8/basic/common"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part8/basic/config"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part8/inventory-srv/handler"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part8/inventory-srv/model"
	proto "github.com/micro-in-cn/tutorials/microservice-in-micro/part8/inventory-srv/proto/inventory"
	tracer "github.com/micro-in-cn/tutorials/microservice-in-micro/part8/plugins/tracer/jaeger"
	"github.com/micro/cli"
	"github.com/micro/go-config/source/grpc"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	openTrace "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
)

var (
	appName = "inv_srv"
	cfg     = &appCfg{}
)

type appCfg struct {
	common.AppCfg
}

func main() {
	// 初始化配置、数据库等信息
	initCfg()

	// 使用consul注册
	micReg := consul.NewRegistry(registryOptions)

	t, io, err := tracer.NewTracer(cfg.Name, "")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)
	// 新建服务
	service := micro.NewService(
		micro.Name(cfg.Name),
		micro.RegisterTTL(time.Second*15),
		micro.RegisterInterval(time.Second*10),
		micro.Registry(micReg),
		micro.Version(cfg.Version),
		micro.WrapHandler(openTrace.NewHandlerWrapper()),
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
