package main

import (
	"context"
	"time"

	proto "github.com/micro-in-cn/tutorials/examples/basic-practices/micro-broker/nsq/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
)

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.broker.kafka.client"),

	)

	srv.Init()

	pub := micro.NewPublisher("go.micro.learning.topic.log", srv.Client())
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
