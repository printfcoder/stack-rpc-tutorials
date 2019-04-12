package basic

import (
	"github.com/micro-in-cn/micro-tutorials/microservice-in-micro/part1/user-service/basic/config"
	"github.com/micro-in-cn/micro-tutorials/microservice-in-micro/part1/user-service/basic/db"
)

func Init() {
	config.InitConfig()
	db.InitDB()
}
