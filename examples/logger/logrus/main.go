package main

import (
	"time"

	"github.com/stack-labs/stack-rpc-plugins/logger/logrus"
	log "github.com/stack-labs/stack-rpc/logger"
)

func main() {
	log.Info("begin")
	logger := logrus.NewLogger(
		log.WithLevel(log.TraceLevel),
		log.Persistence(&log.PersistenceOptions{
			Enable:        true,
			MaxFileSize:   10,
			MaxBackupSize: 500,
			BackupDir:     "/data/var/log/backup",
		}),
		// 将不同级别切成不同文件存储
		logrus.SplitLevel(true),
	)

	log.DefaultLogger = logger

	go func() {
		tk := time.NewTicker(time.Second * 10)
		for {
			select {
			case c := <-tk.C:
				for i := 0; i < 2000; i++ {
					log.Tracef("trace: %s", "hello world!, %v", c.String())
					log.Debugf("debug: %s", "hello world!", c.String())
					log.Info("info: %s", "hello world!", c.String())
					log.Warn("warn: %s", "hello world!", c.String())
					log.Errorf("error: %s", "hello world!", c.String())
				}
			}
		}
	}()

	tk := time.NewTicker(time.Hour * 5)
	select {
	case <-tk.C:
		log.Info("close me")
		return
	}
}
