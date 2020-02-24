package handler

import (
	"context"

	us "github.com/micro-in-cn/tutorials/microservice-in-micro/part1/user-srv/model/user"
	s "github.com/micro-in-cn/tutorials/microservice-in-micro/part1/user-srv/proto/user"
	log "github.com/micro/go-micro/v2/logger"
)

type Service struct{}

var (
	userService us.Service
)

// Init 初始化handler
func Init() {
	var err error
	userService, err = us.GetService()
	if err != nil {
		log.Fatal("[Init] 初始化Handler错误")
		return
	}
}

// QueryUserByName 通过参数中的名字返回用户
func (e *Service) QueryUserByName(ctx context.Context, req *s.Request, rsp *s.Response) error {
	user, err := userService.QueryUserByName(req.UserName)
	if err != nil {
		rsp.Success = false
		rsp.Error = &s.Error{
			Code:   500,
			Detail: err.Error(),
		}

		return nil
	}

	rsp.User = user
	rsp.Success = true
	return nil
}
