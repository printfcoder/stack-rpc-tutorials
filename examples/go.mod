module github.com/stack-labs/stack-rpc-tutorials

go 1.14

replace (
	github.com/stack-labs/stack-rpc v1.0.0 => ../../stack-rpc
	github.com/stack-labs/stack-rpc-plugins/logger/logrus v1.0.0 => ../../stack-rpc-plugins/logger/logrus
)

require (
	github.com/golang/protobuf v1.4.3
	github.com/micro-in-cn/tutorials/examples v0.0.0-20201102112622-908d5548b80c // indirect
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/broker/nats v0.0.0-20200119172437-4fe21aa238fd // indirect
	github.com/micro/go-plugins/broker/nsq v0.0.0-20200119172437-4fe21aa238fd // indirect
	github.com/micro/go-plugins/broker/nsq/v2 v2.9.1 // indirect
	github.com/micro/go-plugins/broker/rabbitmq v0.0.0-20200119172437-4fe21aa238fd // indirect
	github.com/micro/go-plugins/broker/rabbitmq/v2 v2.9.1 // indirect
	github.com/micro/go-plugins/config/source/consul v0.0.0-20200119172437-4fe21aa238fd // indirect
	github.com/micro/go-plugins/config/source/grpc v0.0.0-20200119172437-4fe21aa238fd // indirect
	github.com/micro/go-plugins/config/source/grpc/v2 v2.5.0 // indirect
	github.com/micro/go-plugins/logger/logrus/v2 v2.5.0 // indirect
	github.com/micro/go-plugins/micro/cors/v2 v2.3.0 // indirect
	github.com/micro/go-plugins/registry/consul v0.0.0-20200119172437-4fe21aa238fd // indirect
	github.com/micro/go-plugins/registry/etcdv3 v0.0.0-20200119172437-4fe21aa238fd // indirect
	github.com/micro/micro/v2 v2.3.1 // indirect
	github.com/stack-labs/stack-rpc-plugins/logger/logrus v1.0.0
)
