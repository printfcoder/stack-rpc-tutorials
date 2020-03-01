package main

import (
	"log"

	"github.com/micro-in-cn/tutorials/others/share/learning-go/third-part/greeter/handler"
	"github.com/micro-in-cn/tutorials/others/share/learning-go/third-part/proto/greeter"
	"github.com/micro/go-micro/v2"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.learning.greeter"),
	)

	err := greeter.RegisterGreeterHandler(service.Server(), new(handler.Handler))
	if err != nil {
		log.Fatal(err)
	}

	service.Init()

	log.Fatal(service.Run())
}
