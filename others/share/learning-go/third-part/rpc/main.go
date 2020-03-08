package main

import (
	"log"

	"github.com/micro-in-cn/tutorials/others/share/learning-go/third-part/proto/learning"
	greeterH "github.com/micro-in-cn/tutorials/others/share/learning-go/third-part/rpc/handler/greeter"
	learningH "github.com/micro-in-cn/tutorials/others/share/learning-go/third-part/rpc/handler/learning"
	"github.com/micro/go-micro/v2"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.learning"),
	)

	// 注册Greeter接口
	err := learning.RegisterGreeterHandler(service.Server(), new(greeterH.Handler))
	if err != nil {
		log.Fatal(err)
	}

	// 注册Learning接口
	err = learning.RegisterLearningHandler(service.Server(), new(learningH.Handler))
	if err != nil {
		log.Fatal(err)
	}

	service.Init()

	log.Fatal(service.Run())
}
