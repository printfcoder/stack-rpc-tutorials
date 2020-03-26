package main

import (
	"context"

	proto "github.com/micro-in-cn/tutorials/examples/micro-api/rpc/proto"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/client/grpc"
	"github.com/micro/go-micro/v2/errors"
	log "github.com/micro/go-micro/v2/logger"
)

func main() {
	cli := grpc.NewClient(
		// 根据需要指定重试次数
		client.Retries(4),
		client.Retry(func(ctx context.Context, req client.Request, retryCount int, err error) (b bool, e error) {
			// 遇错重试
			if err != nil {
				// 业务错误
				if err2, ok := err.(*errors.Error); ok {
					// 假设大于1000的都是业务异常
					if err2.Code > 1000 {
						log.Infof("[ERR] 请求错误，业务异常，不重试, err: %s", err)
						return false, nil
					}
				}

				log.Infof("[ERR] 请求错误，第%d次重试，即将重试, err: %s", retryCount, err)
				return true, nil
			}

			return false, nil
		}),
	)

	// 创建客户端
	greeter := proto.NewExampleService("go.micro.retry.example", cli)

	// 调用greeter服务
	for i := 0; i < 10; i++ {
		rsp, err := greeter.Call(context.TODO(), &proto.CallRequest{Name: "Micro中国"})
		if err != nil {
			log.Infof("[ERR] 第%d次 请求发生错误：%s", i, err)
			continue
		}

		log.Infof("[INF] 第%d次 请求结果，%v", i, rsp.Message)
	}
}
