package main

import (
	"time"

	"github.com/stack-labs/stack-rpc"
	"github.com/stack-labs/stack-rpc/config"
	"github.com/stack-labs/stack-rpc/logger"
	"github.com/stack-labs/stack-rpc/plugins/config/source/apollo"
)

func main() {
	service := stack.NewService(stack.ConfigSource(
		apollo.NewSource(
			apollo.Addr("http://127.0.0.1:8080"),
			apollo.Namespaces("application"),
			apollo.AppID("demo"),
			apollo.Cluster("dev"),
		),
	))
	service.Init()

	go func() {
		// 到阿波罗配置后台改变下面的配置值
		for {
			select {
			case <-time.After(1 * time.Second):
				logger.Infof("value: %s", config.Get("demo.server.addr").String("456"))
			}
		}
	}()

	service.Run()
}
