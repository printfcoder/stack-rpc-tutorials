package main

import (
	"github.com/stack-labs/stack-rpc/logger"
)

func main() {
	logger.Init(logger.WithFields(map[string]interface{}{
		"header1": "头1",
	}))

	logger.Info("hello，这条日志带有固定字段")
}
