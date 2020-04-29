package main

import (
	"context"
	"fmt"

	goland "github.com/micro-in-cn/tutorials/others/share/jetbrain/demo/proto"
	"github.com/micro/go-micro/v2"
)

func main() {
	service := micro.NewService(micro.Name("goland.client"))
	service.Init()

	greeter := goland.NewGreeterService("goland.greeter", service.Client())

	rsp, err := greeter.Hello(context.TODO(), &goland.HelloRequest{Name: "Goland 中国"})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Greeting)
}
