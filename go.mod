module github.com/micro-in-cn/tutorials

go 1.13.3

replace (
	github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.0
	github.com/golang/lint => github.com/golang/lint v0.0.0-20190227174305-8f45f776aaf1
	github.com/lucas-clemente/quic-go => github.com/lucas-clemente/quic-go v0.7.1-0.20190913061013-f15a82d3fdc3
	github.com/nats-io/nats.go v1.8.2-0.20190607221125-9f4d16fe7c2d => github.com/nats-io/nats.go v1.8.1
	github.com/nats-io/nkeys => github.com/nats-io/nkeys v0.1.3
	github.com/testcontainers/testcontainer-go => github.com/testcontainers/testcontainers-go v0.0.0-20181115231424-8e868ca12c0f
)

require (
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/alecthomas/units v0.0.0-20190924025748-f65c72e2690d // indirect
	github.com/armon/go-metrics v0.3.0 // indirect
	github.com/chzyer/logex v1.1.10 // indirect
	github.com/chzyer/test v0.0.0-20180213035817-a1ea475d72b1 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-log/log v0.1.0
	github.com/go-redis/redis v6.15.6+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	github.com/gogo/protobuf v1.3.1
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.3.3
	github.com/google/uuid v1.1.1
	github.com/gorilla/mux v1.7.3
	github.com/gorilla/sessions v1.2.0
	github.com/gorilla/websocket v1.4.1
	github.com/grpc-ecosystem/grpc-gateway v1.12.1
	github.com/hashicorp/serf v0.8.5 // indirect
	github.com/klauspost/compress v1.9.1 // indirect
	github.com/klauspost/cpuid v1.2.1 // indirect
	github.com/marten-seemann/qtls v0.4.1 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.1.1-0.20200213185132-d9b3b1758235
	github.com/micro/go-plugins v1.5.1
	github.com/micro/micro/v2 v2.0.0
	github.com/micro/protoc-gen-micro/v2 v2.0.0 // indirect
	github.com/montanaflynn/stats v0.5.0
	github.com/nats-io/nats.go v1.9.1
	github.com/opentracing/opentracing-go v1.1.0
	github.com/pierrec/lz4 v2.3.0+incompatible // indirect
	github.com/prometheus/common v0.7.0
	github.com/rcrowley/go-metrics v0.0.0-20190826022208-cac0b30c2563 // indirect
	github.com/uber/jaeger-client-go v2.20.1+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
	go.uber.org/multierr v1.4.0 // indirect
	go.uber.org/zap v1.13.0
	google.golang.org/appengine v1.6.5 // indirect
	google.golang.org/genproto v0.0.0-20191216164720-4f79533eabd1
	google.golang.org/grpc v1.26.0
	gopkg.in/square/go-jose.v2 v2.4.0 // indirect
	gopkg.in/yaml.v2 v2.2.5 // indirect
)
