package main

import (
	"context"

	proto "github.com/stack-labs/stack-rpc-tutorials/examples/proto/service/rpc"
	"github.com/stack-labs/stack-rpc/client"
	"github.com/stack-labs/stack-rpc/errors"
	log "github.com/stack-labs/stack-rpc/logger"
)

func main() {
	cli := client.NewClient(
		// 根据需要指定重试次数
		client.Retries(4),
		client.Retry(func(ctx context.Context, req client.Request, retryCount int, err error) (b bool, e error) {
			// 遇错重试
			if err != nil {
				// 在这里进行业务代码控制逻辑
				if err2, ok := err.(*errors.Error); ok {
					// 假设大于1000的都是业务异常
					if err2.Code > 1000 {
						log.Infof("[ERR] 请求错误，业务异常，不重试, err: %s", err)
						return false, nil
					}
				}

				log.Infof("[ERR] 请求错误，第%d次重试，即将重试, err: %s", retryCount, err)
				// 返回true，则客户端会进行重试
				return true, nil
			}

			// 没有错误，则返回false，意味不需要重试
			return false, nil
		}),
	)

	// 创建客户端
	greeter := proto.NewGreeterService("stack.rpc.greeter.retry", cli)

	// 调用greeter服务
	for i := 0; i < 10; i++ {
		rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "StackLabs"})
		if err != nil {
			log.Infof("[ERR] 第%d次 请求发生错误：%s", i+1, err)
			continue
		}

		log.Infof("[INF] 第%d次 请求结果，%v", i, rsp.Greeting)
	}
}
