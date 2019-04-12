package db

import (
	"database/sql"
	"fmt"
	"github.com/micro-in-cn/micro-tutorials/microservice-in-micro/part1/user-service/basic/config"
	"log"
	"sync"
)

var (
	inited  bool
	mysqlDB *sql.DB
	m       sync.RWMutex
)

// InitDB 初始化数据库
func InitDB() {
	m.Lock()
	defer m.Unlock()

	var err error

	if inited {
		err = fmt.Errorf("[initMysql] Mysql 已经初始化过")
		log.Fatal(err)
		return
	}

	// 如果配置声明使用mysql
	if config.GetMysqlConfig().GetEnabled() {
		initMysql()
	}

	inited = true
}
