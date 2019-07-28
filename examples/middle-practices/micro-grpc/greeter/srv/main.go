package main

import (
	"context"
	"log"
	"time"

	pb "github.com/micro-in-cn/tutorials/examples/middle-practices/micro-grpc/proto/go/micro"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/service/grpc"
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
	service := grpc.NewService(
		micro.Name("go.micro.srv.greeter"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.Address(":9090"),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	pb.RegisterSayHandler(service.Server(), new(Say))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
