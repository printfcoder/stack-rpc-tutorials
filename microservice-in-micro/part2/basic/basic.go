package basic

import (
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part2/basic/config"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part2/basic/db"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part2/basic/redis"
)

func Init() {
	config.Init()
	db.Init()
	redis.Init()
}
