package main

import (
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"
	"net/http"
	"time"
)

var (
	topic = "go.micro.web.topic.hi"
)

func pub(name string) {
	msg := &broker.Message{
		Header: map[string]string{
			"name": fmt.Sprintf("%s", name),
		},
		Body: []byte(fmt.Sprintf("%s: %s", name, time.Now().String())),
	}
	if err := broker.Publish(topic, msg); err != nil {
		log.Logf("[pub] 发布消息失败： %s", err)
	} else {
		log.Logf("[pub] 发布消息：%s", string(msg.Body))
	}
}

func main() {
	// 创建新服务
	service := web.NewService(
		web.Name("go.micro.book.web.pub"),
		web.Version("latest"),
		web.Address(":8088"),
	)

	// 初始化服务
	_ = service.Init()

	service.HandleFunc("/hi", hi)

	// 运行服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func hi(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	_ = r.ParseForm()
	// 返回结果
	response := map[string]interface{}{
		"ref":  time.Now().UnixNano(),
		"data": "Hello! " + r.Form.Get("name"),
	}

	// 返回JSON结构
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	pub(r.Form.Get("name"))
}
