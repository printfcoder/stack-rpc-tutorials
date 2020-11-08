package main

import (
	"context"
	"encoding/json"
	"strings"

	proto "github.com/micro-in-cn/tutorials/examples/micro-api/api/proto"
	"github.com/micro/go-micro/v2"
	api "github.com/micro/go-micro/v2/api/proto"
	"github.com/micro/go-micro/v2/errors"
	log "github.com/micro/go-micro/v2/logger"
)

type Example struct{}

type Foo struct{}

// Example.Call 通过API向外暴露为/example/call，接收http请求
// 即：/example/call请求会调用go.micro.api.example服务的Example.Call方法
func (e *Example) Call(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Info("Example.Call接口收到请求")

	name, ok := req.Get["name"]

	if !ok || len(name.Values) == 0 {
		return errors.BadRequest("go.micro.api.example", "参数不正确")
	}

	// 打印请求头
	for k, v := range req.Header {
		log.Info("请求头信息，", k, " : ", v)
	}

	rsp.StatusCode = 200

	b, _ := json.Marshal(map[string]string{
		"message": "我们已经收到你的请求，" + strings.Join(name.Values, " "),
	})

	// 设置返回值
	rsp.Body = string(b)

	return nil
}

// Bar 方法全称是Foo.Bar，故而它会以/example/foo/bar为路由提供服务
//
func (f *Foo) Bar(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Info("Foo.Bar接口收到请求")

	if req.Method != "POST" {
		return errors.BadRequest("go.micro.api.example", "require post")
	}

	ct, ok := req.Header["Content-Type"]
	if !ok || len(ct.Values) == 0 {
		return errors.BadRequest("go.micro.api.example", "need content-type")
	}

	if ct.Values[0] != "application/json" {
		return errors.BadRequest("go.micro.api.example", "expect application/json")
	}

	var body map[string]interface{}
	json.Unmarshal([]byte(req.Body), &body)

	// 设置返回值
	rsp.Body = "收到消息：" + string([]byte(req.Body))

	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.example"),
	)

	service.Init()

	// 注册 example handler
	proto.RegisterExampleHandler(service.Server(), new(Example))

	// 注册 foo handler
	proto.RegisterFooHandler(service.Server(), new(Foo))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
