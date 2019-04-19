package handler

import (
	"context"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part3/orders-srv/model/orders"

	proto "github.com/micro-in-cn/tutorials/microservice-in-micro/part3/orders-srv/proto/service"
)

var (
	ordersService orders.Service
)

type Service struct {
}

// Init 初始化handler
func Init() {
	ordersService, _ = orders.GetService()
}

// New 新增订单
func (e *Service) New(ctx context.Context, req *proto.Request, rsp *proto.Response) (err error) {
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
