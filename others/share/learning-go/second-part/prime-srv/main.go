package main

import (
	"github.com/micro-in-cn/tutorials/others/share/learning-go/second-part/prime-srv/handler"
	"github.com/micro-in-cn/tutorials/others/share/learning-go/second-part/proto/prime"
	"github.com/micro/go-micro/v2"
)

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.learning.srv.prime"),
	)

	_ = prime.RegisterPrimeHandler(srv.Server(), handler.Handler())
	srv.Init()

	if err := srv.Run(); err != nil {
		panic(err)
	}
}
