package main

import (
	"context"
	"os"
	"time"

	pb "github.com/micro-in-cn/tutorials/examples/middle-practices/micro-grpc/proto/go/pure-grpc"
	"github.com/micro/go-micro/util/log"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:9090"
	defaultName = "我是来自grpc风格的客户端请求！"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSayClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Hello(ctx, &pb.Request{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Logf("Greeting: %s", r.Msg)
}
