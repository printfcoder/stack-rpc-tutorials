package main

import (
	"github.com/micro/go-log"
	"net/http"

	"github.com/micro-in-cn/tutorials/microservice-in-micro/part3/payment-web/handler"
	"github.com/micro/go-web"
)

func main() {
	// create new web service
	service := web.NewService(
		web.Name("mu.micro.book.web.payment"),
		web.Version("latest"),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// register html handler
	service.Handle("/", http.FileServer(http.Dir("html")))

	// register call handler
	service.HandleFunc("/example/call", handler.ExampleCall)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
