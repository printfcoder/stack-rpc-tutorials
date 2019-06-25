
GOPATH:=$(shell go env GOPATH)

.PHONY: build
build: proto
	GOOS=linux GOARCH=amd64 go build -o user-srv main.go plugin.go

.PHONY: docker
docker:docker
	docker build . -t user-srv:latest