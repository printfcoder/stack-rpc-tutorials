package main

import (
	"log"
	"net/http"

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
	)

	if err := service.Init(); err != nil {
		log.Fatal("Init", err)
	}

	// static files
	service.Handle("/websocket/", http.StripPrefix("/websocket/", http.FileServer(http.Dir("html"))))

	// websocket interface
	service.HandleFunc("/websocket/hi", hi)

	if err := service.Run(); err != nil {
		log.Fatal("Run: ", err)
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
