package main

import (
	"context"
	proto "github.com/micro-in-cn/tutorials/examples/basic-practices/micro-broker/nsq/proto"
	"github.com/micro/go-micro/util/log"
)

var (
	topic = "mu.micro.nsq.demo"

	nsqLookupdAddrs = []string {"172.16.58.20:4161"}
	nsqdAddrs = []string {"172.16.58.20:4150"}
	nsqMaxInFlight = 5
)

type DemoSubscriber struct {}

func (s *DemoSubscriber) Process(ctx context.Context, evt *proto.DemoEvent) error {
	log.Logf("Receive info: Id %d & Timestamp %d\n", evt.Id, evt.Current)
	return nil
}
