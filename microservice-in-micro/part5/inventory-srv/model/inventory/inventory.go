package inventory

import (
	"fmt"
	"sync"
)

var (
	s *service
	m sync.RWMutex
)

// service 服务
type service struct {
}

// Service 库存服务类
type Service interface {
	// Sell 销存
	Sell(bookId, userId int64) (id int64, err error)

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

// Init 初始化库存服务层
func Init() {
	m.Lock()
	defer m.Unlock()

	if s != nil {
		return
	}

	s = &service{}
}
