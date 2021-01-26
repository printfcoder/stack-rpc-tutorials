package main

import (
	"context"
	"fmt"
	"github.com/stack-labs/stack-rpc"
	"time"

	proto "github.com/stack-labs/stack-rpc-tutorials/examples/proto/service/rpc"
	"github.com/stack-labs/stack-rpc/client"
	log "github.com/stack-labs/stack-rpc/logger"
	"github.com/stack-labs/stack-rpc/pkg/metadata"
	"github.com/stack-labs/stack-rpc/registry"
)

func NewCallWrapper1() client.CallWrapper {
	return func(cf client.CallFunc) client.CallFunc {
		return func(ctx context.Context, node *registry.Node, req client.Request, rsp interface{}, opts client.CallOptions) error {
			// WrapCall在发生在均衡之后，此时已经知道要发给哪台服务
			log.Info("[NewCallWrapper1] 准备发往服务节点：", node.Address)

			// 修改请求参数
			body := req.Body().(*proto.HelloRequest)
			body.Name = "StackLabs China"

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
	client := stack.NewService(
		stack.Name("wrap.call.client"),
		stack.WrapCall(NewCallWrapper1(), NewCallWrapper2()),
	)

	client.Init()

	req := client.Client().NewRequest("wrap.call.service", "Greeter.Hello", &proto.HelloRequest{Name: "StackLabs"})
	rsp := &proto.HelloResponse{}

	err := client.Client().Call(context.Background(), req, rsp)
	if err != nil {
		panic(err)
	}

	log.Info(rsp.Greeting)
}
