package main

import (
	b "bytes"
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	pb "github.com/micro-in-cn/tutorials/examples/grpc/proto/go/micro"
	"github.com/micro/go-micro/v2/metadata"
	serverGRPC "github.com/micro/go-micro/v2/server/grpc"
	"github.com/micro/go-micro/v2/service"
	"github.com/micro/go-micro/v2/service/grpc"
)

type Say struct{}

func (s *Say) Hello(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	log.Print("收到请求 Say.Hello 请求")
	md, _ := metadata.FromContext(ctx)
	log.Print("请求头信息 x-user-id: ", md["x-user-id"])
	rsp.Msg = "Hello " + req.Name
	return nil
}

type jsonCodec struct{}

var jsonpbMarshaler = &jsonpb.Marshaler{
	EnumsAsInts:  false,
	EmitDefaults: true,
	OrigName:     true,
}

func (jsonCodec) Marshal(v interface{}) ([]byte, error) {
	if pb, ok := v.(proto.Message); ok {
		s, err := jsonpbMarshaler.MarshalToString(pb)
		return []byte(s), err
	}

	return json.Marshal(v)
}

func (jsonCodec) Unmarshal(data []byte, v interface{}) error {
	if len(data) == 0 {
		return nil
	}
	if pb, ok := v.(proto.Message); ok {
		return jsonpb.Unmarshal(b.NewReader(data), pb)
	}
	return json.Unmarshal(data, v)
}

func (jsonCodec) Name() string {
	return "json"
}

func main() {
	srv := grpc.NewService(
		service.Name("go.micro.srv.greeter"),
		service.RegisterTTL(time.Second*30),
		service.RegisterInterval(time.Second*10),
		service.Address(":9090"),
	)

	// 覆盖掉"application/json"编码器
	srv.Server().Init(serverGRPC.Codec("application/json", &jsonCodec{}))

	srv.Init()

	// Register Handlers
	pb.RegisterSayHandler(srv.Server(), new(Say))

	// Run server
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
