package model

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/micro/go-micro/util/log"
)

var (
	s  *service
	m  sync.RWMutex
	db *sql.DB
)

// service 服务
type service struct {
}

// Init 初始化模型层
func Init() {
	m.Lock()
	defer m.Unlock()

	initMysql()

	if s != nil {
		return
	}
	s = &service{}
}

// Service 库存服务类
type Service interface {
	// Sell 销存
	Sell(bookId, userId int64) (err error)

	// Confirm 确认销存
	Confirm(id int64, state int) (err error)
}

// GetService 获取服务类
func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetService] GetService 未初始化")
	}
	return s, nil
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
