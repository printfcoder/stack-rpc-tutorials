package main

import (
	"github.com/micro/go-plugins/micro/cors/v2"
	"github.com/micro/micro/v2/cmd"
	"github.com/micro/micro/v2/plugin"
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
