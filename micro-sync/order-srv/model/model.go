package model

import (
	"database/sql"
	"fmt"
	"github.com/micro/go-micro"
	"sync"

	invS "github.com/micro-in-cn/tutorials/micro-sync/inventory-srv/proto/inventory"
	proto "github.com/micro-in-cn/tutorials/micro-sync/order-srv/proto/orders"
	"github.com/micro/go-micro/util/log"
)

var (
	s         *service
	invClient invS.InventoryService
	m         sync.RWMutex
	db        *sql.DB
)

// service 服务
type service struct {
}

// Service 订单服务类
type Service interface {
	// New 下单
	New(bookId, userId int64) (orderId int64, err error)

	// GetOrder 获取订单
	GetOrder(orderId int64) (order *proto.Order, err error)

	// UpdateOrderState 更新订单状态
	UpdateOrderState(orderId int64, state int) (err error)
}

// GetService 获取服务类
func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetService] GetService 未初始化")
	}
	return s, nil
}

// Init 初始化模型层
func Init(opts micro.Options) {
	m.Lock()
	defer m.Unlock()

	initMysql()

	if s != nil {
		return
	}
	invClient = invS.NewInventoryService("go.micro.srv.sync.inventory", opts.Client)
	s = &service{}
}

func initMysql() {
	log.Logf("[initMysql] 初始化Mysql")
	var err error

	// 创建连接
	db, err = sql.Open("mysql", "root:123@(127.0.0.1:3306)/micro_book_mall?charset=utf8&parseTime=true&loc=Asia%2FShanghai")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	// 激活链接
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Logf("[initMysql] Mysql 连接成功")
}
