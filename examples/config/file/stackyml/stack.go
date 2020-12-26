package main

import (
	"time"

	"github.com/stack-labs/stack-rpc"
	"github.com/stack-labs/stack-rpc/config"
	log "github.com/stack-labs/stack-rpc/logger"
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
	service.Init()

	log.Infof("demoA: %s", value.IncludeA.DemoA)
	log.Infof("demoB: %s", value.IncludeA.IncludeB.DemoB)
	log.Infof("demoA used get: %s", config.Get("includeA", "demoA").String(""))

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
