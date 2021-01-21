package main

import (
	"context"

	"github.com/stack-labs/stack-rpc"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, name *string, msg *string) error {
	*msg = "Hello " + *name
	return nil
}

func main() {
	// create new service
	service := stack.NewService(
		stack.Name("stack.noproto.greeter"),
	)

	// initialise command line
	service.Init()

	// set the handler
	stack.RegisterHandler(service.Server(), new(Greeter))

	// run service
	service.Run()
}
