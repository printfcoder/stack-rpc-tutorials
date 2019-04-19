package handler

import (
	"context"
	"encoding/json"
	auth "github.com/micro-in-cn/tutorials/microservice-in-micro/part3/auth/proto/auth"
	orders "github.com/micro-in-cn/tutorials/microservice-in-micro/part3/orders-srv/proto/service"
	"github.com/micro/go-log"
	"github.com/micro/go-micro/client"
	"net/http"
	"strconv"
)

var (
	serviceClient orders.Service
	authClient    auth.Service
)

// Error 错误结构体
type Error struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
}

func Init() {
	serviceClient = orders.NewService("mu.micro.book.srv.orders", client.DefaultClient)
	authClient = auth.NewService("mu.micro.book.srv.auth", client.DefaultClient)
}

// New 新增订单入口
func New(w http.ResponseWriter, r *http.Request) {

	// 只接受POST请求
	if r.Method != "POST" {
		log.Logf("非法请求")
		http.Error(w, "非法请求", 400)
		return
	}

	r.ParseForm()

	bookId, _ := strconv.ParseInt(r.Form.Get("userName"), 64, 10)

	// 调用后台服务
	rsp, err := serviceClient.New(context.TODO(), &orders.Request{
		BookId: bookId,
		UserId: 1,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	// 返回JSON结构
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
