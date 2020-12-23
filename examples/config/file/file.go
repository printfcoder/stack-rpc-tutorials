package main

import (
	"github.com/stack-labs/stack-rpc"
	"github.com/stack-labs/stack-rpc/config"
	log "github.com/stack-labs/stack-rpc/logger"
	"time"
)

type includeA struct {
	DemoA    string   `sc:"demoA"`
	IncludeB includeB `sc:"includeB"`
}

type includeB struct {
	DemoB string `sc:"demoB"`
}

type Value struct {
	IncludeA includeA `sc:"includeA"`
}

var (
	value = Value{}
)

func init() {
	config.RegisterOptions(&value)
}

func main() {
	service := stack.NewService()
	service.Init(stack.AfterStart(func() error {
		log.Infof("demoA: %s", value.IncludeA.DemoA)
		log.Infof("demoB: %s", value.IncludeA.IncludeB.DemoB)
		return nil
	}))

	go func() {
		for {
			select {
			case <-time.After(2 * time.Second):
				// try to change DemoB value in includeA.yml
				// there will log the new value
				log.Infof("demoB: %s", value.IncludeA.IncludeB.DemoB)
			}
		}
	}()
	service.Run()
}
