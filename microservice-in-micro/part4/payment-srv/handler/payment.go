package handler

import (
	"context"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part4/payment-srv/model/payment"
	"github.com/micro/go-log"

	proto "github.com/micro-in-cn/tutorials/microservice-in-micro/part4/payment-srv/proto/payment"
)

var (
	paymentService payment.Service
)

type Service struct {
}

// Init 初始化handler
func Init() {
	paymentService, _ = payment.GetService()
}

// New 新增订单
func (e *Service) PayOrder(ctx context.Context, req *proto.Request, rsp *proto.Response) (err error) {
	log.Log("[PayOrder] 收到支付请求")
	err = paymentService.PayOrder(req.OrderId)
	if err != nil {
		rsp.Success = false
		rsp.Error = &proto.Error{
			Detail: err.Error(),
		}
		return
	}

	rsp.Success = true
	return
}
