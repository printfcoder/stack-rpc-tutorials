package main

import (
	"github.com/stack-labs/stack-rpc"
	"github.com/stack-labs/stack-rpc-plugins/service/gateway"
	"github.com/stack-labs/stack-rpc/util/log"
)

func main() {
	svc := stack.NewService(stack.Name("stack.rpc"))
	gateway.Hook(svc)

	// run service
	if err := svc.Run(); err != nil {
		log.Fatal(err)
	}
}
