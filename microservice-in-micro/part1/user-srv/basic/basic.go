package basic

import (
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part1/user-srv/basic/config"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part1/user-srv/basic/db"
)

func Init() {
	config.Init()
	db.Init()
}
