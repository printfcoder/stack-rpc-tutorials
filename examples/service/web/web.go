package main

import (
	"net/http"

	"github.com/stack-labs/stack-rpc"
	"github.com/stack-labs/stack-rpc/service/web"
)

func main() {
	s := stack.NewWebService(
		stack.Name("stack.rpc.web"),
		stack.Address(":8080"),
		stack.WebRootPath("web-demo"),
		stack.WebHandleFuncs(
			web.HandlerFunc{
				Route: "hello",
				Func: func(w http.ResponseWriter, r *http.Request) {
					w.Write([]byte(`hello world`))
				},
			},
		),
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
