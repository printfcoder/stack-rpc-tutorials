package main

import (
	"net/http"

	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"
)

func main() {
	service := web.NewService(
		// 服务名，可填可不填，建议作为必填
		web.Name("go.micro.book.web.pub"),
		// 以下选项都是可选的！
		web.Version("latest"),
		web.Address(":8088"),
		web.Id("123"),
	)

	err := service.Init(
		web.BeforeStart(func() error {
			log.Error("[web] 启动前的动作执行了")
			return nil
		}),
		web.AfterStart(func() error {
			log.Error("[web] 启动后的动作执行了")
			return nil
		}))
	if err != nil {
		panic(err)
	}

	// 静态文件
	// static files
	service.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("html"))))

	err = service.Run()
	if err != nil {
		panic(err)
	}
}
