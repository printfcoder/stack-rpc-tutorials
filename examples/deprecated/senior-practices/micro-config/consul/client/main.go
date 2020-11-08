package main

import (
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/config/source/consul"
)

type Appconfig struct {
	Address string `json:"address"`
	Port    int64  `json:"port"`
}

func main() {
	consulSource := consul.NewSource(
		consul.WithAddress("127.0.0.1:8500"),
		consul.WithPrefix("/micro/app"),
	)

	conf := config.NewConfig()

	err := conf.Load(consulSource)
	if err != nil {
		log.Fatal(err)
	}

	var Conf Appconfig
	Conf = Appconfig{}

	err = conf.Get("micro", "app", "mysql").Scan(&Conf)
	if err != nil {
		log.Fatal(err)
	}

	log.Logf("Read conf: %s", Conf.Port)

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
