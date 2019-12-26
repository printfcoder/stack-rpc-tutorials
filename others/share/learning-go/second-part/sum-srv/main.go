package main

import (
	"context"
	"github.com/micro/cli"

	logProto "github.com/micro-in-cn/tutorials/others/share/learning-go/second-part/proto/log"
	"github.com/micro-in-cn/tutorials/others/share/learning-go/second-part/proto/sum"
	"github.com/micro-in-cn/tutorials/others/share/learning-go/second-part/sum-srv/handler"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/util/log"
	_ "github.com/micro/go-plugins/broker/rabbitmq"
)

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.learning.srv.sum"),
		micro.WrapHandler(
			// 并行请求只支持5个
			rateLimiter(5),
		),
		micro.Flags(cli.StringFlag{
			Name:   "learning_go",
			EnvVar: "LEARNING_GO",
			Usage:  "help一下，你就知道",
		}),
	)

	srv.Init(
		micro.WrapHandler(
			reqLogger(srv.Client()),
		),
		micro.BeforeStart(func() error {
			log.Error("[srv] 启动前的动作执行了")
			return nil
		}),
		micro.AfterStart(func() error {
			log.Error("[srv] 启动后的动作执行了")
			return nil
		}),
	)

	_ = sum.RegisterSumHandler(srv.Server(), handler.Handler())

	if err := srv.Run(); err != nil {
		panic(err)
	}
}

func rateLimiter(rate int) server.HandlerWrapper {
	tokens := make(chan bool, rate)
	for i := 0; i < rate; i++ {
		tokens <- true
	}

	return func(handler server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			token := <-tokens
			defer func() {
				tokens <- token
			}()
			return handler(ctx, req, rsp)
		}
	}
}

func reqLogger(cli client.Client) server.HandlerWrapper {
	pub := micro.NewPublisher("go.micro.learning.topic.log", cli)

	return func(handler server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			log.Log("发送日志")
			evt := logProto.LogEvt{
				Msg: "Hello",
			}

			_ = pub.Publish(context.TODO(), &evt)
			return handler(ctx, req, rsp)
		}
	}
}
