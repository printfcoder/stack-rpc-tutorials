package model

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/micro-in-cn/tutorials/micro-sync/common"
	invS "github.com/micro-in-cn/tutorials/micro-sync/inventory-srv/proto/inventory"
	ordS "github.com/micro-in-cn/tutorials/micro-sync/order-srv/proto/orders"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
)

var (
	db           *sql.DB
	s            *service
	invClient    invS.InventoryService
	ordSClient   ordS.OrdersService
	m            sync.RWMutex
	payPublisher micro.Publisher
)

// Init 初始化模型层
func Init(opts micro.Options) {
	m.Lock()
	defer m.Unlock()

	initMysql()

	if s != nil {
		return
	}

	invClient = invS.NewInventoryService("go.micro.sync.srv.inventory", opts.Client)
	ordSClient = ordS.NewOrdersService("go.micro.sync.srv.orders", opts.Client)
	payPublisher = micro.NewPublisher(common.TopicPaymentDone, opts.Client)
	s = &service{}
}

func initMysql() {
	log.Logf("[initMysql] 初始化Mysql")
	// 创建连接
	mysqlDB, err := sql.Open("mysql", "root:123@(127.0.0.1:3306)/micro_book_mall?charset=utf8&parseTime=true&loc=Asia%2FShanghai")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	// 激活链接
	if err = mysqlDB.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Logf("[initMysql] Mysql 连接成功")
}

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
