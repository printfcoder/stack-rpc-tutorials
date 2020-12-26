package stack

import (
	"time"

	"github.com/stack-labs/stack-rpc"
	stack2 "github.com/stack-labs/stack-rpc-plugins/config/source/stack"
	"github.com/stack-labs/stack-rpc/config"
	"github.com/stack-labs/stack-rpc/logger"
)

func main() {
	service := stack.NewService(
		stack.New(
			stack2.NewSource(),
		),
	)

	service.Init()

	go func() {
		// 到阿波罗配置后台改变下面的配置值
		for {
			select {
			case <-time.After(1 * time.Second):
				logger.Infof("value: %s", config.Get("demo.server.addr").String("456"))
			}
		}
	}()

	service.Run()
}
