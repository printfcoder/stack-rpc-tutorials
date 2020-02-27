package main

import (
	"context"
	"net/http"
	"strconv"

	"github.com/micro-in-cn/tutorials/others/share/learning-go/second-part/proto/sum"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/web"
)

var (
	srvClient sum.SumService
)

func main() {
	service := web.NewService(
		// 服务名，可填可不填，建议作为必填
		web.Name("go.micro.learning.web.portal"),
		// 以下选项都是可选的！
		web.Version("latest"),
		web.Address(":8088"),
		web.StaticDir("html"),
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

	srvClient = sum.NewSumService("go.micro.learning.srv.sum", service.Options().Service.Client())

	// 静态文件
	// static files
	service.Handle("/learning", http.StripPrefix("/learning", http.FileServer(http.Dir("html"))))
	service.HandleFunc("/learning/sum", Sum)
	err = service.Run()
	if err != nil {
		panic(err)
	}
}

func Sum(w http.ResponseWriter, r *http.Request) {
	inputStr := r.URL.Query().Get("input")
	input, _ := strconv.ParseInt(inputStr, 10, 10)
	req := &sum.SumRequest{
		Input: input,
	}

	rsp, err := srvClient.GetSum(context.Background(), req)
	if err != nil {
		// todo
	}

	w.Write([]byte(strconv.Itoa(int(rsp.Output))))
}
