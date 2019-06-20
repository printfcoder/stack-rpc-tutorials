module github.com/micro-in-cn/tutorials

go 1.12

replace github.com/testcontainers/testcontainer-go => github.com/testcontainers/testcontainers-go v0.0.0-20181115231424-8e868ca12c0f

replace github.com/golang/lint => github.com/golang/lint v0.0.0-20190227174305-8f45f776aaf1

require (
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-redis/redis v6.15.2+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/protobuf v1.3.1
	github.com/google/uuid v1.1.1
	github.com/gorilla/mux v1.7.1
	github.com/gorilla/sessions v1.1.3
	github.com/micro/cli v0.1.0
	github.com/micro/go-api v0.7.0
	github.com/micro/go-config v1.1.0
	github.com/micro/go-log v0.1.0
	github.com/micro/go-micro v1.1.0
	github.com/micro/go-plugins v1.1.0
	github.com/micro/go-web v1.0.0
	github.com/micro/micro v1.1.1
	github.com/opentracing/opentracing-go v1.1.0
	github.com/uber/jaeger-client-go v2.16.0+incompatible
	go.uber.org/zap v1.9.1
	google.golang.org/grpc v1.19.1
)
