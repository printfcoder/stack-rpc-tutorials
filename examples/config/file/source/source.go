package main

import (
	"time"

	"github.com/stack-labs/stack-rpc"
	"github.com/stack-labs/stack-rpc/config"
	log "github.com/stack-labs/stack-rpc/logger"
	"github.com/stack-labs/stack-rpc/pkg/config/source/file"
)

type source struct {
	DemoA string `sc:"demoA"`
}

type Value struct {
	Source source `sc:"source"`
}

var (
	value Value
)

func init() {
	config.RegisterOptions(&value)
}

func main() {
	service := stack.NewService(
		stack.Config(config.NewConfig(config.Source(file.NewSource(file.WithPath("./source.yml"))))),
	)
	service.Init()

	log.Infof("demoA: %s", value.Source.DemoA)

	go func() {
		for {
			select {
			case <-time.After(2 * time.Second):
				// try to change DemoA value in source.yml
				// there will log the new value
				log.Infof("demoA: %s", value.Source.DemoA)
			}
		}
	}()
	service.Run()
}
