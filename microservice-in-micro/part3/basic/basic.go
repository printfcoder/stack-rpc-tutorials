package basic

import (
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part3/basic/config"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part3/basic/db"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part3/basic/redis"
)

func Init() {
	config.Init()
	db.Init()
	redis.Init()
}
