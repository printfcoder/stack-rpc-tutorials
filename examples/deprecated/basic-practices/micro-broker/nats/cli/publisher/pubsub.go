package publisher

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	pubsub "github.com/micro-in-cn/tutorials/examples/basic-practices/micro-broker/nats/srv/proto/pubsub"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/util/log"
)

func SendEvent(topic string, b broker.Broker) {
	t := time.NewTicker(5 * time.Second)

	var i int
	for _ = range t.C {
		ev := pubsub.Event{
			Id:        uuid.New().String(),
			Timestamp: time.Now().Unix(),
			Message:   fmt.Sprintf("如果你看到了消息 %s, '那是因为我一直爱着你", topic),
		}

		body, _ := json.Marshal(ev)
		msg := &broker.Message{
			Header: map[string]string{
				"id": fmt.Sprintf("%d", i),
			},
			Body: body,
		}

		log.Logf("publishing %+v", ev)

		if err := b.Publish(topic, msg); err != nil {
			log.Logf("error publishing: %v", err)
		}
		i++
	}
}
