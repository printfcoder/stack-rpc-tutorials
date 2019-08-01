package publisher

import (
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro/util/log"
	"github.com/pborman/uuid"
	"time"
	"github.com/micro/go-micro/broker"

	pubsub "github.com/brucewangno1/go-micro-pubsub-with-nats/srv/proto/pubsub"
)

func SendEvent(topic string, b broker.Broker) {
	t := time.NewTicker(5*time.Second)

	var i int
	for _ = range t.C {
		ev := pubsub.Event{
			Id: uuid.NewUUID().String(),
			Timestamp: time.Now().Unix(),
			Message: fmt.Sprintf("Message you all day on %s, 'cause nothing's gonna change my love for you", topic),
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