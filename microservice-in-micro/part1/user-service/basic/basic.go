package basic

import (
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part1/user-service/basic/config"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part1/user-service/basic/db"
)

func Init() {
	config.Init()
	db.Init()
}
