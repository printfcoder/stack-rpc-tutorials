package main

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/registry/mdns"
)

// 请求参数结构只要对方服务能识别的就行
type whatEverReq struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

// 响应结构只要我方服务能识别的就行
type whatEverRsp struct {
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func main() {
	cli := client.NewClient(
		// 与目录服务同注册中心即可
		client.Registry(mdns.NewRegistry()),
	)

	// 调用目标服务的结构
	req := cli.NewRequest("go.micro.rpc.example", "Example.Call",
		&whatEverReq{
			Name: "John",
		},
		// 不确定对方服务时，需要使用JSON格式，而不是protobuf
		client.WithContentType("application/json"))

	// 自定义元数据
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id": "john",
		"X-From-Id": "script",
	})

	rsp := &whatEverRsp{}

	// 调用服务
	if err := cli.Call(ctx, req, rsp); err != nil {
		fmt.Println("call err: ", err, rsp)
		return
	}

	fmt.Println("rsp: ", rsp.Message)
}
