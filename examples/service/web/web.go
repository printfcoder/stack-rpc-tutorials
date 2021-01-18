package main

import (
	"net/http"

	"github.com/stack-labs/stack-rpc"
	"github.com/stack-labs/stack-rpc/service/web"
)

func main() {
	s := stack.NewWebService(
		// 服务名
		stack.Name("stack.rpc.web"),
		// 地址与端口
		stack.Address(":8080"),
		// 根目录
		stack.WebRootPath("web-demo"),
		// 路由与Handler
		stack.WebHandleFuncs(
			web.HandlerFunc{
				Route: "hello",
				Func: func(w http.ResponseWriter, r *http.Request) {
					w.Write([]byte(`hello world`))
				},
			},
		),
		// 静态目录
		stack.WebStaticDir("webapp", "static"),
	)
	err := s.Init()
	if err != nil {
		panic(err)
	}

	err = s.Run()
	if err != nil {
		panic(err)
	}
}
