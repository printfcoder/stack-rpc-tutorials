package main

import (
	"github.com/stack-labs/stack-rpc"
	log "github.com/stack-labs/stack-rpc/logger"
)

func main() {
	service := stack.NewService()
	service.Init()
	log.Debug("hello，这是Debug级别")
	log.Info("hello，这是Info级别")
}
