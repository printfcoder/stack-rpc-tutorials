package main

import (
	"context"
	"fmt"
	"testing"

	proto "github.com/micro-in-cn/tutorials/examples/service/proto"
	"github.com/micro/go-micro/v2"
)

var (
	greeter proto.GreeterService
)

func init() {
	service := micro.NewService(micro.Name("greeter.client"))
	service.Init()
	greeter = proto.NewGreeterService("greeter.service", service.Client())
}

func TestHello(t *testing.T) {
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "Micro中国"})
	if err != nil {
		t.Error(err)
	}

	if rsp.Greeting == "" {
		t.Error(fmt.Errorf("[ERR] invalid rsp"))
	}

	t.Log("[INFO] rsp:", rsp.Greeting)
}
