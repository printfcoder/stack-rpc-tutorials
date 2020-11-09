package main

import (
	"github.com/stack-labs/stack-rpc/config"
	"github.com/stack-labs/stack-rpc/logger"
	"github.com/stack-labs/stack-rpc/plugins/config/source/apollo"
)

func main() {
	c, err := config.NewConfig(
		config.Storage(true),
	)

	if err != nil {
		logger.Fatal(err)
	}

	err = c.Load(
		apollo.NewSource(
			apollo.Addr("http://127.0.0.1:8080"),
			apollo.Namespaces("application"),
			apollo.AppID("demo-app"),
			apollo.Cluster("dev"),
		))
	if err != nil {
		logger.Error(err)
	}

	logger.Info(c.Get("test.config1").String("test.config1 default value"))
	logger.Info(c.Get("test.config2").String("test.config1 default value"))
}
