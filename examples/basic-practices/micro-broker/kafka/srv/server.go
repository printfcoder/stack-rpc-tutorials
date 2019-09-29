package main

import (
	"context"

	proto "github.com/micro-in-cn/tutorials/examples/basic-practices/micro-broker/nsq/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/broker/kafka"
)

type Sub struct {
}

func (s *Sub) Process(ctx context.Context, evt *proto.DemoEvent) error {
	log.Logf("Receive info: Id %d & Timestamp %d\n", evt.Id, evt.Current)
	return nil
}

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.broker.kafka.srv"),
		micro.Broker(kafka.NewBroker(
			broker.Addrs([]string{"127.0.0.1:9092"}...),
		)),
	)

	srv.Init()

	_ = micro.RegisterSubscriber("go.micro.broker.topic.kafka", srv.Server(), &Sub{})

	if err := srv.Run(); err != nil {
		log.Fatalf("error occurs: %v", err)
	}
}
