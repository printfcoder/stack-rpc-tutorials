package main

import (
	"context"
	"log"
	"time"

	pb "github.com/micro-in-cn/tutorials/examples/grpc/proto/go/micro"
	"github.com/micro/go-micro/v2/metadata"
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

func main() {
	srv := grpc.NewService(
		service.Name("go.micro.srv.greeter"),
		service.RegisterTTL(time.Second*30),
		service.RegisterInterval(time.Second*10),
		service.Address(":9090"),
	)

	// optionally setup command line usage
	srv.Init()

	// Register Handlers
	pb.RegisterSayHandler(srv.Server(), new(Say))

	// Run server
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
