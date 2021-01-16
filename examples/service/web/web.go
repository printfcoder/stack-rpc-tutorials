package main

import (
	"net/http"

	"github.com/stack-labs/stack-rpc/service"
	"github.com/stack-labs/stack-rpc/service/web"
)

func main() {
	s := web.NewService(
		service.Name("stack.rpc.web"),
		service.Address(":8080"),
		web.RootPath("web-demo"),
		web.HandleFuncs(
			web.HandlerFunc{
				Route: "hello",
				Func: func(w http.ResponseWriter, r *http.Request) {
					w.Write([]byte(`hello world`))
				},
			},
		),
		web.StaticDir("webapp", "static"),
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
