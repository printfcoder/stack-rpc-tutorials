package main

import (
	"context"

	proto "github.com/micro-in-cn/tutorials/others/share/learning-go/second-part/proto/log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	_ "github.com/micro/go-plugins/broker/rabbitmq"
)

type Sub struct {
}

func (s *Sub) Process(ctx context.Context, evt *proto.LogEvt) error {
	log.Logf("收到日志 %v\n", evt.Msg)
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.learning.srv.log"),
		micro.Version("latest"),
	)

	service.Init()

	_ = micro.RegisterSubscriber("go.micro.learning.topic.log", service.Server(), &Sub{})

	if err := service.Run(); err != nil {
		panic(err)
	}
}
