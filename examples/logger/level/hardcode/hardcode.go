package main

import log "github.com/stack-labs/stack-rpc/logger"

func main() {
	log.Init(log.WithLevel(log.DebugLevel))
	log.Debug("hello，这是Debug级别")

	log.Init(log.WithLevel(log.InfoLevel))
	log.Debug("hello，这是Debug级别")
	log.Info("hello，这是Info级别")
}
