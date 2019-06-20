package access

import (
	"fmt"
	"sync"

	r "github.com/go-redis/redis"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part5/basic/config"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part5/plugins/jwt"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part5/plugins/redis"
	"github.com/micro/go-micro/util/log"
)

var (
	s   *service
	ca  *r.Client
	m   sync.RWMutex
	cfg = &jwt.Jwt{}
)

// service 服务
type service struct {
}

// Service 用户服务类
type Service interface {
	// MakeAccessToken 生成token
	MakeAccessToken(subject *Subject) (ret string, err error)

	// GetCachedAccessToken 获取缓存的token
	GetCachedAccessToken(subject *Subject) (ret string, err error)

	// DelUserAccessToken 清除用户token
	DelUserAccessToken(token string) (err error)
}

// GetService 获取服务类
func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetService] GetService 未初始化")
	}
	return s, nil
}

// Init 初始化用户服务层
func Init() {
	m.Lock()
	defer m.Unlock()

	if s != nil {
		return
	}

	err := config.C().App("jwt", cfg)
	if err != nil {
		panic(err)
	}

	log.Logf("[initCfg] 配置，cfg：%v", cfg)

	ca = redis.Redis()

	s = &service{}
}
