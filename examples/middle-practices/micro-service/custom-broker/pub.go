package main

import (
	"fmt"
	"log"
	"time"

	"github.com/micro/go-micro/broker"
)

func main() {

	go func() {
		tick := time.NewTicker(time.Second)
		i := 0
		for range tick.C {
			// 定义第一条消息
			msg1 := &broker.Message{
				Header: map[string]string{
					"id": fmt.Sprintf("%d", i),
				},
				Body: []byte(fmt.Sprintf("主题1消息%d: %s", i, time.Now().String())),
			}
			if err := broker.Publish("go.micro.topic.custom1", msg1); err != nil {
				log.Printf("[pub] 发布消息1失败： %v", err)
			} else {
				fmt.Println("[pub] 发布消息1：", string(msg1.Body))
			}

			// 定义第二条消息
			msg2 := &broker.Message{
				Header: map[string]string{
					"id": fmt.Sprintf("%d", i),
				},
				Body: []byte(fmt.Sprintf("主题2消息%d: %s", i, time.Now().String())),
			}
			if err := broker.Publish("go.micro.topic.custom2", msg2); err != nil {
				log.Printf("[pub] 发布消息2失败： %v", err)
			} else {
				fmt.Println("[pub] 发布消息2：", string(msg2.Body))
			}

			i++
		}
	}()

	<-time.After(10 * time.Second)
}
