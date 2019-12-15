package main

import (
	"context"
	"encoding/json"
	logProto "github.com/micro-in-cn/tutorials/others/share/learning-go/second-part/proto/log"
	"github.com/micro-in-cn/tutorials/others/share/learning-go/second-part/proto/sum"
	"github.com/micro-in-cn/tutorials/others/share/learning-go/second-part/sum-srv/handler"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/util/log"
)

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.learning.srv.sum"),
		micro.WrapHandler(
			// 并行请求只支持5个
			rateLimiter(5),
			reqLogger(),
		),
	)

	srv.Init()

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

func reqLogger() server.HandlerWrapper {
	return func(handler server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			log.Log("发送日志")
			evt := logProto.LogEvt{
				Msg: "Hello",
			}

			body, _ := json.Marshal(evt)
			broker.Publish("go.micro.learning.topic.log",
				&broker.Message{
					Header: map[string]string{
						"serviceName": "sum",
					},
					Body: body,
				})
			return handler(ctx, req, rsp)
		}
	}
}
