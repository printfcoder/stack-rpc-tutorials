
GOPATH:=$(shell go env GOPATH)

.PHONY:
build: proto
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o payment-web main.go plugin.go

.PHONY: docker
docker:docker
	docker build . -t payment-web:latest
