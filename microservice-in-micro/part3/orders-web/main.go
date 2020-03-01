package main

import (
	"fmt"
	"net/http"

	"github.com/micro-in-cn/tutorials/microservice-in-micro/part3/basic"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part3/basic/config"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part3/orders-web/handler"
	"github.com/micro/cli/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
)

func main() {
	// 初始化配置
	basic.Init()

	// 使用etcd注册
	micReg := etcd.NewRegistry(registryOptions)

	// 创建新服务
	service := web.NewService(
		web.Name("mu.micro.book.web.orders"),
		web.Version("latest"),
		web.Registry(micReg),
		web.Address(":8091"),
	)

	// 初始化服务
	if err := service.Init(
		web.Action(
			func(c *cli.Context) {
				// 初始化handler
				handler.Init()
			}),
	); err != nil {
		log.Fatal(err)
	}

	// 新建订单接口
	authHandler := http.HandlerFunc(handler.New)
	service.Handle("/orders/new", handler.AuthWrapper(authHandler))
	service.HandleFunc("/", handler.Hello)

	// 运行服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}
