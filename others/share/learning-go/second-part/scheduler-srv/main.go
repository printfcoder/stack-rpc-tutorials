package main

import "github.com/micro/go-micro"

func main() {
	srv := micro.NewService(micro.Name("go.micro.srv.learning"))

	srv.Init()

	if err := srv.Run(); err != nil {
		panic(err)
	}
}
