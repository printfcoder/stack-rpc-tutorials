package main

import (
	"log"
	"time"

	"context"
	"github.com/micro/go-micro"
)

func main() {

	// 获取上下文，并声明关闭函数cancel
	ctx, cancel := context.WithCancel(context.Background())

	// 5秒后关闭
	go func() {
		<-time.After(time.Second * 5)
		log.Println("关闭示例: 收到关闭信号")
		cancel()
	}()

	// 创建服务
	service := micro.NewService(
		// 一定要使用上面声明的ctx
		micro.Context(ctx),
	)

	// 初始化服务
	service.Init()

	// 运行服务
	service.Run()
}
