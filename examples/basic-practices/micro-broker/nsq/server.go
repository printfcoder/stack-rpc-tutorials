package main

import (
	"github.com/micro-in-cn/tutorials/examples/basic-practices/micro-broker/nsq/config"
	"github.com/micro-in-cn/tutorials/examples/basic-practices/micro-broker/nsq/pubsub"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/broker/nsq"
)

func main() {
	srv := micro.NewService(
		micro.Name("mu.micro.srv.demo"),
		micro.Broker(nsq.NewBroker(
			nsq.WithLookupdAddrs(config.NsqLookupdAddrs),
			broker.Addrs(config.NsqdAddrs...),
		)),
	)

	srv.Init()

	sOpts := broker.NewSubscribeOptions(
		nsq.WithMaxInFlight(config.NsqMaxInFlight),
	)

	_ = micro.RegisterSubscriber(config.Topic, srv.Server(), &pubsub.DemoSubscriber{}, server.SubscriberContext(sOpts.Context))

	if err := srv.Run(); err != nil {
		log.Fatalf("error occurs: %v", err)
	}
}
