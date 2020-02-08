package main

import (
	"github.com/micro/go-plugins/micro/cors"
	"github.com/micro/micro/cmd"
	"github.com/micro/micro/plugin"
)

func init() {
	// 注册跨域插件
	if err := plugin.Register(cors.NewPlugin()); err != nil {
		panic(err)
	}
}

func main() {
	cmd.Init()
}
