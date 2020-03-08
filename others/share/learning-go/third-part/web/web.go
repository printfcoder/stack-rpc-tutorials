package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/transport/grpc"
	"github.com/micro/go-micro/v2/web"
)

func main() {
	// New web service
	service := web.NewService(
		web.Name("go.micro.web.learning"),
		web.MicroService(micro.NewService(micro.Transport(grpc.NewTransport()))),
	)

	service.Options().Service.Client()

	if err := service.Init(); err != nil {
		log.Fatal("Init", err)
	}

	service.HandleFunc("/web/hi", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")

		_ = r.ParseForm()
		// 返回结果
		response := map[string]interface{}{
			"ref":  time.Now().UnixNano(),
			"data": fmt.Sprintf("Hello! %s. Welcome to Learning Web!", r.Form.Get("name")),
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	})

	if err := service.Run(); err != nil {
		log.Fatal("Run: ", err)
	}
}
