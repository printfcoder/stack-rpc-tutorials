package subscriber

import (
	"context"
	"encoding/json"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part3/orders-srv/model/orders"
	payS "github.com/micro-in-cn/tutorials/microservice-in-micro/part3/payment-srv/proto/service"
	"github.com/micro/go-log"
	"github.com/micro/go-micro/broker"
)

var (
	ordersService orders.Service
)

// Init 初始化handler
func Init() {
	ordersService, _ = orders.GetService()
}

// PayOrder 订单支付消息
func PayOrder(ctx context.Context, msg *broker.Message) (err error) {

	var payment payS.Payments
	err = json.Unmarshal(msg.Body, payment)
	if err != nil {
		log.Logf("[PayOrder] 解序列化消息失败，err：%s", err)
		return
	}

	log.Logf("[PayOrder] 收到支付订单通知，%d，%d", payment.OrderId, payment.State)
	err = ordersService.UpdateOrderState(payment.OrderId, int(payment.State))
	if err != nil {
		log.Logf("[PayOrder] 收到支付订单通知，更新状态异常，%s", err)
		return
	}
	return
}
