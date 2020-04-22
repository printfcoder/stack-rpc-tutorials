package main

import (
	"context"

	proto "github.com/micro-in-cn/tutorials/others/share/learning-go/second-part/proto/log"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	_ "github.com/micro/go-plugins/broker/rabbitmq/v2"
)

type Sub struct {
}

func (s *Sub) Process(ctx context.Context, evt *proto.LogEvt) error {
	log.Infof("收到日志 %v\n", evt.Msg)
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
