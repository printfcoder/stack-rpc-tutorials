package handler

import (
	"context"

	"github.com/micro-in-cn/tutorials/microservice-in-micro/part4/orders-srv/model/orders"
	proto "github.com/micro-in-cn/tutorials/microservice-in-micro/part4/orders-srv/proto/orders"
	log "github.com/micro/go-micro/v2/logger"
)

var (
	ordersService orders.Service
)

type Orders struct {
}

// Init 初始化handler
func Init() {
	ordersService, _ = orders.GetService()
}

// New 新增订单
func (e *Orders) New(ctx context.Context, req *proto.Request, rsp *proto.Response) (err error) {
	orderId, err := ordersService.New(req.BookId, req.UserId)
	if err != nil {
		rsp.Success = false
		rsp.Error = &proto.Error{
			Detail: err.Error(),
		}
		return
	}

	rsp.Order = &proto.Order{
		Id: orderId,
	}
	return
}

// GetOrder 获取订单
func (e *Orders) GetOrder(ctx context.Context, req *proto.Request, rsp *proto.Response) (err error) {
	log.Infof("[GetOrder] 收到获取订单请求，%d", req.OrderId)

	rsp.Order, err = ordersService.GetOrder(req.OrderId)
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
