package main

import (
	"context"
	"fmt"
	"time"

	"github.com/stack-labs/stack-rpc"
	proto "github.com/stack-labs/stack-rpc-tutorials/examples/proto/service/stream/msg"
	"github.com/stack-labs/stack-rpc/client"
	log "github.com/stack-labs/stack-rpc/logger"
)

var msgService proto.MsgStreamService

var c client.Client

func SendMsg() {
	next, _ := c.Options().Selector.Select("msg.service")
	node, _ := next()
	stream, err := msgService.Msg(context.Background(), func(options *client.CallOptions) {
		// 指定节点
		options.Address = []string{node.Address}
	})

	go func() {
		id := 1
		for {
			err = stream.Send(&proto.StreamMessage{
				Namespace: "stack.stream.msg.demo.client",
				Id:        "any-id-we-want-" + fmt.Sprint(id),
				Fields: []*proto.Field{
					{
						Name: "id",
						Type: "string",
					},
					{
						Name: "Name",
						Type: "string",
					},
				},
				Payload: []*proto.Payload{
					{
						Value: []string{fmt.Sprint(id), "name-" + fmt.Sprint(id)},
					},
				},
			})

			id++
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			time.Sleep(5 * time.Second)

			ret, err := msgService.HeartBeat(context.Background(), &proto.HeartBeatRequest{
				AreYouOk: "Client1",
			}, func(options *client.CallOptions) {
				// 指定节点
				options.Address = []string{node.Address}
			})

			if err != nil {
				log.Errorf("send heartbeat err: %s", err)
			}

			log.Infof("got heart back: %v", ret)
		}
	}()

	select {
	case <-time.After(time.Second * 20):
		return
	}
}

func main() {
	// 定义服务，可以传入其它可选参数
	service := stack.NewService(stack.Name("msg.client"))
	service.Init()

	// 创建客户端
	c = service.Client()
	msgService = proto.NewMsgStreamService("msg.service", c)
	SendMsg()
}
