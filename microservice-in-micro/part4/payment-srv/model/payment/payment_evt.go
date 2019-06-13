package payment

import (
	"context"
	"github.com/google/uuid"
	proto "github.com/micro-in-cn/tutorials/microservice-in-micro/part4/payment-srv/proto/payment"
	"github.com/micro/go-log"
	"time"
)

// sendPayDoneEvt 发送支付事件
func (s *service) sendPayDoneEvt(orderId int64, state int32) {
	// 构建事件
	ev := &proto.PayEvent{
		Id:       uuid.New().String(),
		SentTime: time.Now().Unix(),
		OrderId:  orderId,
		State:    state,
	}

	log.Logf("[sendPayDoneEvt] 发送支付事件，%+v\n", ev)

	// 广播
	if err := payPublisher.Publish(context.Background(), ev); err != nil {
		log.Logf("[sendPayDoneEvt] 异常: %v", err)
	}
}
