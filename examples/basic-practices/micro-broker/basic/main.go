package main

import (
	"fmt"
	"time"

	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/config/cmd"
	"github.com/micro/go-micro/util/log"
)

var (
	topic = "mu.micro.book.topic.payment.done"
	b broker.Broker
)

func pub() {
	tick := time.NewTicker(time.Second)
	i := 0
	for range tick.C {
		msg := &broker.Message{
			Header: map[string]string{
				"id": fmt.Sprintf("%d", i),
			},
			Body: []byte(fmt.Sprintf("%d: %s", i, time.Now().String())),
		}
		log.Infof(broker.String())
		if err := broker.Publish(topic, msg); err != nil {
			log.Infof("[pub] Message publication failed: %v", err)
		} else {
			fmt.Println("[pub] Message published: ", string(msg.Body))
		}
		i++
	}
}

func sub() {
	_, err := broker.Subscribe(topic, func(p broker.Event) error {
		log.Infof("[sub] Received Body: %s, Header: %s\n", string(p.Message().Body), p.Message().Header)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	// cmd.Init() parses flags and env variables.
	// If you leave out cmd.Init(),
	// broker "http" will be used as default
	// other than ones like nats you have specified.
	cmd.Init()
	if err := broker.Init(); err != nil {
		log.Fatalf("broker.Init() error: %v", err)
	}
	if err := broker.Connect(); err != nil {
		log.Fatalf("broker.Connect() error: %v", err)
	}

	go pub()
	go sub()

	<-time.After(time.Second * 20)
}
