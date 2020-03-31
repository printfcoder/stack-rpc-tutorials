package main

import (
	"context"

	proto "github.com/micro-in-cn/tutorials/examples/basic-practices/micro-broker/nsq/proto"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
)

type Sub struct {
}

func (s *Sub) Process(ctx context.Context, evt *proto.DemoEvent) error {
	log.Infof("Receive info: Id %d & Timestamp %d\n", evt.Id, evt.Current)
	return nil
}

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.broker.kafka.srv"),
	)

	srv.Init()

	_ = micro.RegisterSubscriber("go.micro.learning.topic.log", srv.Server(), &Sub{})

	if err := srv.Run(); err != nil {
		log.Fatalf("error occurs: %v", err)
	}
}
