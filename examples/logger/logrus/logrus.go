package main

import (
	"time"

	"github.com/stack-labs/stack-rpc"
	"github.com/stack-labs/stack-rpc-plugins/logger/logrus"
	log "github.com/stack-labs/stack-rpc/logger"
)

func main() {
	service := stack.NewService(
		stack.Logger(logrus.NewLogger(
			log.WithLevel(log.TraceLevel),
			log.Persistence(&log.PersistenceOptions{
				Enable:                true,
				MaxFileSize:           10,
				MaxBackupSize:         500,
				MaxBackupKeepDays:     1,
				FileNamePattern:       "",
				BackupFileNamePattern: "",
				Dir:                   "/tmp/logs",
				BackupDir:             "/tmp/logs/backup",
			}),
			// 将不同级别切成不同文件存储
			logrus.SplitLevel(true),
		)))

	service.Init()

	go func() {
		tk := time.NewTicker(time.Second * 2)
		for {
			select {
			case c := <-tk.C:
				for i := 0; i < 10; i++ {
					log.Tracef("hello world! %s", c.String())
					log.Debugf("hello world!  %s", c.String())
					log.Infof("hello world! %s", c.String())
					log.Warnf("hello world! %s", c.String())
					log.Errorf("hello world! %s", c.String())
				}
			}
		}
	}()

	service.Run()
}
