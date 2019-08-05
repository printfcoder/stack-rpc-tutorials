package main

import (
	"context"
	"github.com/micro-in-cn/tutorials/examples/basic-practices/micro-broker/nsq/config"
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
			nsq.WithLookupdAddrs(config.NsqLookupdAddrs),
			broker.Addrs(config.NsqdAddrs...),
		)),
	)

	srv.Init()

	pub := micro.NewPublisher(config.Topic, srv.Client())
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
