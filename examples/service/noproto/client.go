package main

import (
	"context"
	"fmt"

	"github.com/stack-labs/stack-rpc"
	"github.com/stack-labs/stack-rpc/client"
)

func main() {
	service := stack.NewService()
	service.Init()
	c := service.Client()

	request := c.NewRequest("stack.noproto.greeter", "Greeter.Hello", "StackLabs", client.WithContentType("application/json"))
	var response string

	if err := c.Call(context.TODO(), request, &response); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response)
}
