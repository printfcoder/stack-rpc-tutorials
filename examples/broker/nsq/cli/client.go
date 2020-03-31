package main

import (
	"context"
	"time"

	proto "github.com/micro-in-cn/tutorials/examples/basic-practices/micro-broker/nsq/proto"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/broker/nsq/v2"
)

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.broker.nsq.client"),
		micro.Broker(nsq.NewBroker(
			broker.Addrs([]string{"127.0.0.1:4150"}...),
		)),
	)

	srv.Init()

	pub := micro.NewEvent("go.micro.broker.topic.nsq", srv.Client())
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)

			_ = pub.Publish(context.TODO(), &proto.DemoEvent{
				Id:      int32(i),
				Current: time.Now().Unix(),
			})
		}
	}()

	if err := srv.Run(); err != nil {
		log.Fatalf("error occurs: %v", err)
	}
}
