package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/broker/nsq"
	"time"
)

func main() {
	srv := micro.NewService(
		micro.Name("mu.micro.srv.demo"),
		micro.Broker(nsq.NewBroker(
			nsq.WithLookupdAddrs(nsqLookupdAddrs),
			broker.Addrs(nsqdAddrs...),
		)),
		micro.RegisterInterval(10 * time.Second),
	)

	srv.Init()

	sOpts := broker.NewSubscribeOptions(
		nsq.WithMaxInFlight(nsqMaxInFlight),
	)

	_ = micro.RegisterSubscriber(topic, srv.Server(), &DemoSubscriber{}, server.SubscriberContext(sOpts.Context))

	if err := srv.Run(); err != nil {
		log.Fatalf("error occurs: %v", err)
	}
}
