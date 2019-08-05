package main

import (
	"context"
	proto "github.com/micro-in-cn/tutorials/examples/basic-practices/micro-broker/nsq/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/broker/nsq"
	"time"
)

func main() {
	srv := micro.NewService(
		micro.Name("mu.micro.cli.demo"),
		micro.Broker(nsq.NewBroker(
			nsq.WithLookupdAddrs(nsqLookupdAddrs),
			broker.Addrs(nsqdAddrs...),
		)),
	)

	srv.Init()

	pub := micro.NewPublisher(topic, srv.Client())
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)

			_ = pub.Publish(context.TODO(), &proto.DemoEvent{
				Id: int32(i),
				Current: time.Now().Unix(),
			})
		}
	}()

	if err := srv.Run(); err != nil {
		log.Fatalf("error occurs: %v", err)
	}
}
