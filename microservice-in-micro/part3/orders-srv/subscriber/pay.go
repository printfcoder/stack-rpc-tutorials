package subscriber

import (
	"context"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part3/orders-srv/model/orders"
	payS "github.com/micro-in-cn/tutorials/microservice-in-micro/part3/payment-srv/proto/service"
	"github.com/micro/go-log"
)

var (
	ordersService orders.Service
)

// Init 初始化handler
func Init() {
	ordersService, _ = orders.GetService()
}

// PayOrder 订单支付消息
func PayOrder(ctx context.Context, msg *payS.Payments) (err error) {
	log.Logf("[PayOrder] 收到支付订单通知，%d，%d", msg.OrderId, msg.State)
	err = ordersService.UpdateOrderState(msg.OrderId, int(msg.State))
	if err != nil {
		log.Logf("[PayOrder] 收到支付订单通知，更新状态异常，%s", err)
		return
	}
	return
}
