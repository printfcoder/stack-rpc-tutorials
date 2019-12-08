package main

import (
	"github.com/micro-in-cn/tutorials/others/share/learning-go/second-part/proto/sum"
	"github.com/micro-in-cn/tutorials/others/share/learning-go/second-part/sum-srv/handler"
	"github.com/micro/go-micro"
)

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.learning.srv.sum"),
		micro.Address("127.0.0.1:9988"),
	)

	_ = sum.RegisterSumHandler(srv.Server(), handler.Handler())
	srv.Init()

	if err := srv.Run(); err != nil {
		panic(err)
	}
}
