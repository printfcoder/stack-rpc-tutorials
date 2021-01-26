package main

import (
	"github.com/stack-labs/stack-rpc"
	"github.com/stack-labs/stack-rpc/broker"
)

func main() {
	sub := stack.NewService(stack.Name("stack.broker.sub"))
	sub.Init()

	s, _ := broker.Subscribe("stack.broker.demo.topic", func(event broker.Event) error {
		return nil
	})

	broker.Publish()

	sub.Run()
}
