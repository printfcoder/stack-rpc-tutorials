package main

import (
	"context"
	"encoding/json"

	proto "github.com/micro-in-cn/tutorials/examples/micro-api/rpc/proto"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
)

type Example struct{}

func (e *Example) Call(ctx context.Context, req *proto.CallRequest, rsp *proto.CallResponse) error {
	log.Info("Example.Call接口收到请求，返回成功")

	b, _ := json.Marshal(map[string]string{
		"message": "我们已经收到你的请求，" + req.Name,
	})

	// 设置返回值
	rsp.Message = string(b)

	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.retry.example"),
	)

	service.Init()

	// 注册 example handler
	proto.RegisterExampleHandler(service.Server(), new(Example))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
