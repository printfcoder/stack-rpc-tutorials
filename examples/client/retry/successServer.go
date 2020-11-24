package main

import (
	"context"
	"encoding/json"

	"github.com/stack-labs/stack-rpc"
	proto "github.com/stack-labs/stack-rpc-tutorials/examples/proto/service/rpc"
	log "github.com/stack-labs/stack-rpc/logger"
)

type SuccessExample struct{}

func (e *SuccessExample) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	log.Info("SuccessExample.Hello，返回成功")

	b, _ := json.Marshal(map[string]string{
		"message": "我们已经收到你的请求，" + req.Name,
	})

	// 设置返回值
	rsp.Greeting = string(b)

	return nil
}

func main() {
	service := stack.NewService(
		stack.Name("stack.rpc.greeter.retry"),
	)

	service.Init()

	// 注册 example handler
	proto.RegisterGreeterHandler(service.Server(), new(SuccessExample))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
