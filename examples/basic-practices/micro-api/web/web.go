package main

import (
	"encoding/json"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/transport/grpc"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/micro/go-micro/web"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func main() {
	// New web service
	service := web.NewService(
		web.Name("go.micro.web.websocket"),
		web.MicroService(micro.NewService(micro.Transport(grpc.NewTransport()))),
	)

	service.Options().Service.Client()

	if err := service.Init(); err != nil {
		log.Fatal("Init", err)
	}

	// static files
	service.Handle("/websocket/", http.StripPrefix("/websocket/", http.FileServer(http.Dir("html"))))

	// websocket interface
	service.HandleFunc("/websocket/hi", hi)

	// websocket interface
	service.HandleFunc("/websocket/hi2", hi2)

	// websocket interface
	service.HandleFunc("/websocket/hi3/hi3", hi2)

	if err := service.Run(); err != nil {
		log.Fatal("Run: ", err)
	}
}

func hi2(w http.ResponseWriter, r *http.Request) {
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
}

func hi(w http.ResponseWriter, r *http.Request) {

	c, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("upgrade: %s", err)
		return
	}

	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		log.Printf("recv: %s", message)

		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
