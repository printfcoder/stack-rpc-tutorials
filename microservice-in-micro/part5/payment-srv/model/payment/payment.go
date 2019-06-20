package payment

import (
	"fmt"
	"sync"

	"github.com/micro-in-cn/tutorials/microservice-in-micro/part5/basic/common"
	invS "github.com/micro-in-cn/tutorials/microservice-in-micro/part5/inventory-srv/proto/inventory"
	ordS "github.com/micro-in-cn/tutorials/microservice-in-micro/part5/orders-srv/proto/orders"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
)

var (
	s            *service
	invClient    invS.InventoryService
	ordSClient   ordS.OrdersService
	m            sync.RWMutex
	payPublisher micro.Publisher
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

	invClient = invS.NewInventoryService("mu.micro.book.srv.inventory", client.DefaultClient)
	ordSClient = ordS.NewOrdersService("mu.micro.book.srv.orders", client.DefaultClient)
	payPublisher = micro.NewPublisher(common.TopicPaymentDone, client.DefaultClient)
	s = &service{}
}
