package main

import (
	"context"
	"fmt"

	"github.com/stack-labs/stack-rpc"
	proto "github.com/stack-labs/stack-rpc-tutorials/examples/proto/service/stream/msg"
	"github.com/stack-labs/stack-rpc/errors"
	log "github.com/stack-labs/stack-rpc/logger"
)

type MsgStreamHandler struct{}

func (m *MsgStreamHandler) Msg(ctx context.Context, msg proto.MsgStream_MsgStream) error {
	for {
		b, err := msg.Recv()
		if err != nil {
			return errors.InternalServerError("msg.service", err.Error())
		}

		log.Infof("id: %v; fields: %v; namespace: %v; extends: %v; payload: %v", b.Id, b.Fields, b.Namespace, b.Extends, b.Payload)

		msg.SendMsg(&proto.StreamResponse{
			Status: "OK",
		})
	}
}

func (m *MsgStreamHandler) HeartBeat(ctx context.Context, req *proto.HeartBeatRequest, rsp *proto.HeartBeatRespond) error {
	rsp.ImOK = "OK！ " + req.AreYouOk

	return nil
}

func main() {
	// 创建服务，除了服务名，其它选项可加可不加，比如Version版本号、Metadata元数据等
	service := stack.NewService(
		stack.Name("msg.service"),
		stack.Version("latest"),
	)
	service.Init()

	// 注册服务
	_ = proto.RegisterMsgStreamHandler(service.Server(), new(MsgStreamHandler))

	// 启动服务
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
