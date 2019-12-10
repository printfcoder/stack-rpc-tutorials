package main

import (
	"context"
	"github.com/micro-in-cn/tutorials/others/share/learning-go/second-part/proto/sum"
	"github.com/micro-in-cn/tutorials/others/share/learning-go/second-part/sum-srv/handler"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/util/log"
)

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.learning.srv.sum"),
		// 并发只支持5次
		micro.WrapHandler(rateLimiter(5)),
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
				// 处理完成后释放
				log.Infof("[rateLimiter] 释放限制")
				tokens <- token
			}()
			return handler(ctx, req, rsp)
		}
	}
}
