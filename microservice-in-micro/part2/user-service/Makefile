
GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/service/service.proto

.PHONY: build
build: proto

	go build -o user-service main.go plugin.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t user-service:latest
