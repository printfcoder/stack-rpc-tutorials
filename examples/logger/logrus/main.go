package main

import (
	"github.com/stack-labs/stack-rpc-plugins/logger/logrus"
	log "github.com/stack-labs/stack-rpc/logger"
)

func main() {
	log.Info("begin")
	logger := logrus.NewLogger(
		log.WithLevel(log.TraceLevel),
		log.Persistence(&log.PersistenceOptions{
			Enable:        true,
			MaxFileSize:   1,
			MaxBackupSize: 10,
		}),
		// 将不同级别切成不同文件存储
		logrus.SplitLevel(true),
	)

	log.DefaultLogger = logger

	for i := 0; i < 20000; i++ {
		log.Tracef("trace: %s", "hello world!")
		log.Debugf("debug: %s", "hello world!")
		log.Info("info: %s", "hello world!")
		log.Warn("warn: %s", "hello world!")
		log.Errorf("error: %s", "hello world!")
	}
}
