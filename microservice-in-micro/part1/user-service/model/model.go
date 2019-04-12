package model

import "github.com/micro-in-cn/micro-tutorials/microservice-in-micro/part1/user-service/model/user"

// InitModel 初始化模型层
func InitModel() {
	user.InitUserService()
}
