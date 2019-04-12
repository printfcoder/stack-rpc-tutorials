package user

import (
	"fmt"
	proto "github.com/micro-in-cn/micro-tutorials/microservice-in-micro/part1/user-service/proto/service"
	"sync"
)

// service 服务
type service struct {
}

var (
	s *service
	m sync.RWMutex
)

// Service 用户服务类
type Service interface {
	// QueryUserByName 根据用户名获取用户
	QueryUserByName(userName string) (ret *proto.User, err error)
}

// GetService 获取服务类
func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetService] GetService 未初始化")
	}
	return s, nil
}

// InitUserService 初始化用户服务层
func InitUserService() {
	m.Lock()
	defer m.Unlock()

	if s != nil {
		return
	}

	s = &service{}
}
