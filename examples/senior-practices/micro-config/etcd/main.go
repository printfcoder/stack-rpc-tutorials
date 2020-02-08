package main

import (
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/etcd"
	"github.com/micro/go-micro/util/log"
)

type Micro struct {
	Config Config `json:"config"`
}

type Config struct {
	Mysql MysqlConfig `json:"mysql"`
}

type MysqlConfig struct {
	Port int `json:"port"`
}

func main() {
	etcdSource := etcd.NewSource(
		etcd.WithAddress("127.0.0.1:2379"),
		etcd.WithPrefix("/micro/app"),
	)

	conf := config.NewConfig()
	err := conf.Load(etcdSource)
	if err != nil {
		log.Fatal(err)
	}

	microCfg := &struct {
		Micro Micro `json:"micro"`
	}{}

	v := conf.Get("micro", "app", "config")
	err = v.Scan(&microCfg)
	if err != nil {
		log.Fatal(err)
	}

	log.Logf("Read mysql port: %d", microCfg.Micro.Config.Mysql.Port)

	//开始侦听变动事件
	watcher, err := conf.Watch()
	if err != nil {
		log.Fatal(err)
	}

	log.Logf("Watch changes ...")
	for {
		v, err := watcher.Next()
		if err != nil {
			log.Fatal(err)
		}

		log.Logf("Watch changes: %v", string(v.Bytes()))
	}
}
