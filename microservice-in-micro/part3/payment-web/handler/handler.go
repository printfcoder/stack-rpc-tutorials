package handler

import (
	"context"
	"encoding/json"
	auth "github.com/micro-in-cn/tutorials/microservice-in-micro/part3/auth/proto/auth"
	payS "github.com/micro-in-cn/tutorials/microservice-in-micro/part3/payment-srv/proto/service"
	"github.com/micro/go-log"
	"github.com/micro/go-micro/client"
	"net/http"
	"strconv"
	"time"
)

var (
	serviceClient payS.Service
	authClient    auth.Service
)

// Error 错误结构体
type Error struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
}

func Init() {
	serviceClient = payS.NewService("mu.micro.book.srv.payment", client.DefaultClient)
	authClient = auth.NewService("mu.micro.book.srv.auth", client.DefaultClient)
}

// PayOrder 支付订单
func PayOrder(w http.ResponseWriter, r *http.Request) {

	// 只接受POST请求
	if r.Method != "POST" {
		log.Logf("非法请求")
		http.Error(w, "非法请求", 400)
		return
	}

	r.ParseForm()

	orderId, _ := strconv.ParseInt(r.Form.Get("orderId"), 64, 10)

	// 调用后台服务
	rsp, err := serviceClient.PayOrder(context.TODO(), &payS.Request{
		OrderId: orderId,
	})
	if err != nil {
		log.Logf("[PayOrder] 支付失败，err：%s", err)
		http.Error(w, err.Error(), 500)
		return
	}

	// 返回结果
	response := map[string]interface{}{
		"success": rsp.Success,
		"ref":     time.Now().UnixNano(),
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	// 返回JSON结构
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
