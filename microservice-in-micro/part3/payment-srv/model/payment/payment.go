package payment

import (
	"fmt"
	invS "github.com/micro-in-cn/tutorials/microservice-in-micro/part3/inventory-srv/proto/service"
	ordS "github.com/micro-in-cn/tutorials/microservice-in-micro/part3/orders-srv/proto/service"
	"github.com/micro/go-grpc/client"
	"sync"
)

var (
	s          *service
	invClient  invS.Service
	ordSClient ordS.Service
	m          sync.RWMutex
)

// service 服务
type service struct {
}

// Service 服务类
type Service interface {
	// PayOrder 支付订单
	PayOrder(orderId int64) (err error)
}

// GetService 获取服务类
func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetService] GetService 未初始化")
	}
	return s, nil
}

// Init 初始化库存服务层
func Init() {
	m.Lock()
	defer m.Unlock()

	if s != nil {
		return
	}

	invClient = invS.NewService("mu.micro.book.srv.inventory", client.DefaultClient)
	ordSClient = ordS.NewService("mu.micro.book.srv.orders", client.DefaultClient)
	s = &service{}
}
