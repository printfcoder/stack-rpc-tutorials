module github.com/stack-labs/stack-rpc-tutorials/examples

go 1.14

replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.4
	github.com/stack-labs/stack-rpc v1.0.0 => ../../stack-rpc
	github.com/stack-labs/stack-rpc-plugins/config/source/apollo v1.0.0 => ../../stack-rpc-plugins/config/source/apollo
	github.com/stack-labs/stack-rpc-plugins/logger/logrus v1.0.0 => ../../stack-rpc-plugins/logger/logrus
	github.com/stack-labs/stack-rpc-plugins/registry/consul v1.0.0 => ../../stack-rpc-plugins/registry/consul
	github.com/stack-labs/stack-rpc-plugins/service/stackway v1.0.0 => ../../stack-rpc-plugins/service/stackway
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)

require (
	github.com/beevik/ntp v0.3.0 // indirect
	github.com/caddyserver/certmagic v0.10.6 // indirect
	github.com/coreos/bbolt v1.3.4 // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/fsouza/go-dockerclient v1.6.5 // indirect
	github.com/go-git/go-git/v5 v5.1.0 // indirect
	github.com/gobwas/httphead v0.0.0-20180130184737-2c6c146eadee // indirect
	github.com/gobwas/pool v0.2.0 // indirect
	github.com/gobwas/ws v1.0.3 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.3
	github.com/google/go-cmp v0.5.0 // indirect
	github.com/google/uuid v1.1.1
	github.com/gorilla/websocket v1.4.2
	github.com/grpc-ecosystem/grpc-gateway v1.9.5
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/lib/pq v1.8.0 // indirect
	github.com/magiconair/properties v1.8.4 // indirect
	github.com/marten-seemann/qtls v0.10.0 // indirect
	github.com/marten-seemann/qtls-go1-15 v0.1.1 // indirect
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.4.0
	github.com/micro/go-plugins/broker/nats v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/broker/nsq v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/broker/nsq/v2 v2.3.0
	github.com/micro/go-plugins/broker/rabbitmq v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/broker/rabbitmq/v2 v2.3.0
	github.com/micro/go-plugins/config/source/consul v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/config/source/grpc v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/config/source/grpc/v2 v2.3.0 // indirect
	github.com/micro/go-plugins/logger/logrus/v2 v2.3.0
	github.com/micro/go-plugins/micro/cors/v2 v2.3.0
	github.com/micro/go-plugins/registry/consul v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/registry/etcdv3 v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/micro/v2 v2.3.1
	github.com/nats-io/jwt v1.1.0 // indirect
	github.com/nats-io/nats.go v1.10.0 // indirect
	github.com/prometheus/common v0.6.0
	github.com/spf13/viper v1.7.1 // indirect
	github.com/stack-labs/stack-rpc v1.0.0
	github.com/stack-labs/stack-rpc-plugins/config/source/apollo v1.0.0
	github.com/stack-labs/stack-rpc-plugins/logger/logrus v1.0.0
	github.com/stack-labs/stack-rpc-plugins/service/stackway v1.0.0
	github.com/stack-labs/stack-rpc-plugins/registry/consul v1.0.0
	github.com/tevid/gohamcrest v1.1.1 // indirect
	github.com/tmc/grpc-websocket-proxy v0.0.0-20200427203606-3cfed13b9966 // indirect
	go.etcd.io/bbolt v1.3.5 // indirect
	go.opencensus.io v0.22.2 // indirect
	golang.org/x/crypto v0.0.0-20201016220609-9e8e0b390897 // indirect
	golang.org/x/exp v0.0.0-20200331195152-e8c3332aa8e5 // indirect
	golang.org/x/net v0.0.0-20200707034311-ab3426394381 // indirect
	golang.org/x/sys v0.0.0-20201107080550-4d91cf3a1aaf // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.25.0
	gopkg.in/urfave/cli.v1 v1.20.0 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)
