package main

import (
	"math/rand"
	"time"

	"github.com/stack-labs/stack-rpc"
	"github.com/stack-labs/stack-rpc-plugins/logger/logrus"
	log "github.com/stack-labs/stack-rpc/logger"
)

func main() {
	service := stack.NewService(
		stack.Name("stack.rpc.logger"),
		stack.Logger(
			logrus.NewLogger(
				log.WithLevel(log.TraceLevel),
				// 将不同级别切成不同文件存储
				logrus.SplitLevel(false),
			)))

	err := service.Init()
	if err != nil {
		panic(err)
	}
	rand.Seed(time.Now().UnixNano())
	go func() {
		rTime := time.Duration(rand.Int31n(1000)) * time.Millisecond
		tk := time.NewTicker(rTime)
		for {
			select {
			case c := <-tk.C:
				log.Tracef("I'm Tracef. hello world! %s", c.String())
				log.Debugf("I'm Debugf. hello world!  %s", c.String())
				log.Infof("I'm Infof. hello world! %s", c.String())
				log.Warnf("I'm Warnf. hello world! %s", c.String())
				log.Errorf("I'm Errorf. hello world! %s", c.String())
			}
		}
	}()

	err = service.Run()
	if err != nil {
		panic(err)
	}
}
