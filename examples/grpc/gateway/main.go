package main

import (
	"context"
	"flag"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pb "github.com/micro-in-cn/tutorials/examples/grpc/proto/go/grpc-gw/greeter"
	log "github.com/micro/go-micro/v2/logger"
	"google.golang.org/grpc"
)

var (
	// the go.micro.srv.greeter address
	endpoint = flag.String("endpoint", "localhost:9090", "go.micro.srv.greeter address")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := pb.RegisterSayHandlerFromEndpoint(ctx, mux, *endpoint, opts)
	if err != nil {
		return err
	}

	log.Info("[grpc gateway] Listening on 8080")

	return http.ListenAndServe(":8080", mux)
}

func main() {
	flag.Parse()

	if err := run(); err != nil {
		log.Fatal(err)
	}
}
