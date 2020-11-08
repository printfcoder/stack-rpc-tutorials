package main

import (
	"github.com/micro/go-micro/v2/logger"
)

func main() {
	logger.Init(logger.WithLevel(logger.DebugLevel))

	logger.Debug("Debug")
	logger.Debugf("Debug %s", "Hello")

	logger.Init(logger.WithLevel(logger.InfoLevel))
	logger.Debug("Debug2")
	logger.Debugf("Debug2 %s", "Hello")

	logger.Info("Info")
	logger.Infof("Info %s", "Hello")
	logger.Error("Error")
	logger.Errorf("Debug %s", "Hello")
}
