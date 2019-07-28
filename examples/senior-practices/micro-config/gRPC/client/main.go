package main

import (
	"github.com/micro/go-config"
	grpcConfig "github.com/micro/go-config/source/grpc"
	"github.com/micro/go-log"
)

type Micro struct {
	Info
}

type Info struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
	Message string `json:"message,omitempty"`
	Age     int    `json:"age,omitempty"`
}

func main() {

	// 声明使用grpc配置中心
	source := grpcConfig.NewSource(
		grpcConfig.WithAddress("127.0.0.1:9600"),
		grpcConfig.WithPath("/micro"),
	)
	conf := config.NewConfig()

	// 加载配置
	err := conf.Load(source)
	if err != nil {
		log.Fatal(err)
	}

	// 渲染配置到指定结构
	configs := &Micro{}
	err = conf.Scan(configs)
	if err != nil {
		log.Fatal(err)
	}

	log.Logf("Read configs: %s", string(conf.Bytes()))

	// 开始侦听变动事件
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
