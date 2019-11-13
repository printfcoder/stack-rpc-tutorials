package main

import (
	"context"
	"fmt"
	"time"

	proto "github.com/micro-in-cn/tutorials/examples/basic-practices/micro-service/proto"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/util/log"
)

func NewCallWrapper1() client.CallWrapper {
	return func(cf client.CallFunc) client.CallFunc {
		return func(ctx context.Context, node *registry.Node, req client.Request, rsp interface{}, opts client.CallOptions) error {
			// WrapCall在发生在均衡之后，此时已经知道要发给哪台服务
			log.Info("[NewCallWrapper1] 准备发往服务节点：", node.Address)

			// 修改请求参数
			body := req.Body().(*proto.HelloRequest)
			body.Name = "Micro China"

			// 模拟随机错误
			if time.Now().Second() < 20 {
				return fmt.Errorf("随机错误，请求中止！")
			}

			newMd := metadata.Metadata{}
			newMd["call-wrapped1"] = "call-wrapped-value1"
			ctx = metadata.NewContext(ctx, newMd)
			err := cf(ctx, node, req, rsp, opts)
			return err
		}
	}
}

func NewCallWrapper2() client.CallWrapper {
	return func(cf client.CallFunc) client.CallFunc {
		return func(ctx context.Context, node *registry.Node, req client.Request, rsp interface{}, opts client.CallOptions) error {
			log.Info("[NewCallWrapper2] Wrapper工作：", node.Address)

			newMd, b := metadata.FromContext(ctx)
			if !b {
				newMd = metadata.Metadata{}
			}
			newMd["call-wrapped2"] = "call-wrapped-value2"
			ctx = metadata.NewContext(ctx, newMd)
			err := cf(ctx, node, req, rsp, opts)
			return err
		}
	}
}

func main() {
	cli := client.NewClient(
		client.WrapCall(
			NewCallWrapper1(),
			NewCallWrapper2()),
	)

	req := client.NewRequest("wrap.call.service", "Greeter.Hello", &proto.HelloRequest{Name: "Micro中国"})
	rsp := &proto.HelloResponse{}

	err := cli.Call(context.Background(), req, rsp)
	if err != nil {
		panic(err)
	}

	log.Info(rsp.Greeting)
}
