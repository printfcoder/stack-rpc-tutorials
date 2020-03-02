package db

import (
	"database/sql"

	"github.com/micro-in-cn/tutorials/microservice-in-micro/part4/basic/config"
	log "github.com/micro/go-micro/v2/logger"
)

type db struct {
	Mysql Mysql `json："mysql"`
}

// Mysql mySQL 配置
type Mysql struct {
	URL               string `json:"url"`
	Enable            bool   `json:"enabled"`
	MaxIdleConnection int    `json:"maxIdleConnection"`
	MaxOpenConnection int    `json:"maxOpenConnection"`
}

func initMysql() {
	log.Info("[initMysql] 初始化Mysql")

	c := config.C()
	cfg := &db{}

	err := c.App("db", cfg)
	if err != nil {
		log.Warnf("[initMysql] %s", err)
	}

	if !cfg.Mysql.Enable {
		log.Warnf("[initMysql] 未启用Mysql")
		return
	}

	// 创建连接
	mysqlDB, err = sql.Open("mysql", cfg.Mysql.URL)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	// 最大连接数
	mysqlDB.SetMaxOpenConns(cfg.Mysql.MaxOpenConnection)

	// 最大闲置数
	mysqlDB.SetMaxIdleConns(cfg.Mysql.MaxIdleConnection)

	// 激活链接
	if err = mysqlDB.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Warnf("[initMysql] Mysql 连接成功")
}
