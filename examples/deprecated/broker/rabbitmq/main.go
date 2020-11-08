package main

import (
	"fmt"
	"log"
	"time"

	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/config/cmd"
)

var (
	topic = "mu.micro.book.topic.payment.done"
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
		if err := broker.Publish(topic, msg); err != nil {
			log.Printf("[pub] 发布消息失败： %v", err)
		} else {
			fmt.Println("[pub] 发布消息：", string(msg.Body))
		}
		i++
	}
}

func sub() {
	_, err := broker.Subscribe(topic, func(p broker.Event) error {
		fmt.Printf("[sub] Received Body: %s, Header: %s", string(p.Message().Body), p.Message().Header)
		return nil
	}, broker.Queue("mu.micro.book.topic.queue"))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	cmd.Init()

	if err := broker.Init(); err != nil {
		log.Fatalf("Broker 初始化错误：%v", err)
	}
	if err := broker.Connect(); err != nil {
		log.Fatalf("Broker 连接错误：%v", err)
	}

	go pub()
	go sub()

	<-time.After(time.Second * 100)
}
