package main

import (
	"context"

	proto "github.com/micro-in-cn/tutorials/examples/basic-practices/micro-broker/nsq/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/broker/nsq"
)

type Sub struct {
}

func (s *Sub) Process(ctx context.Context, evt *proto.DemoEvent) error {
	log.Logf("Receive info: Id %d & Timestamp %d\n", evt.Id, evt.Current)
	return nil
}

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.broker.nsq.srv"),
		micro.Broker(nsq.NewBroker(
			broker.Addrs([]string{"127.0.0.1:4150"}...),
		)),
	)

	srv.Init()

	sOpts := broker.NewSubscribeOptions(
		nsq.WithMaxInFlight(5),
	)

	_ = micro.RegisterSubscriber("go.micro.broker.topic.nsq", srv.Server(), &Sub{}, server.SubscriberContext(sOpts.Context))

	if err := srv.Run(); err != nil {
		log.Fatalf("error occurs: %v", err)
	}
}
