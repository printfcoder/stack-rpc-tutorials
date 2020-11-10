module github.com/stack-labs/stack-rpc-tutorials/examples

go 1.14

replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.4
	github.com/stack-labs/stack-rpc v1.0.0 => ../../stack-rpc
	github.com/stack-labs/stack-rpc-plugins/logger/logrus v1.0.0 => ../../stack-rpc-plugins/logger/logrus
	github.com/stack-labs/stack-rpc/plugins/config/source/apollo v1.0.0 => ../../stack-rpc-plugins/config/source/apollo
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)

require (
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.1.1
	github.com/gorilla/websocket v1.4.2
	github.com/grpc-ecosystem/grpc-gateway v1.9.5
	github.com/micro-in-cn/tutorials/examples v0.0.0-20201102112622-908d5548b80c
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/broker/nats v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/broker/nsq v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/broker/nsq/v2 v2.9.1
	github.com/micro/go-plugins/broker/rabbitmq v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/broker/rabbitmq/v2 v2.9.1
	github.com/micro/go-plugins/config/source/consul v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/config/source/grpc v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/config/source/grpc/v2 v2.5.0 // indirect
	github.com/micro/go-plugins/logger/logrus/v2 v2.5.0
	github.com/micro/go-plugins/micro/cors/v2 v2.3.0
	github.com/micro/go-plugins/registry/consul v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/registry/etcdv3 v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/micro/v2 v2.3.1
	github.com/prometheus/common v0.6.0
	github.com/stack-labs/stack-rpc v1.0.0
	github.com/stack-labs/stack-rpc-plugins/logger/logrus v1.0.0
	github.com/stack-labs/stack-rpc/plugins/config/source/apollo v1.0.0
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013
	google.golang.org/grpc v1.29.1
)
