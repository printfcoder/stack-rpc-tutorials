module github.com/micro-in-cn/tutorials

go 1.12

replace github.com/testcontainers/testcontainer-go => github.com/testcontainers/testcontainers-go v0.0.0-20181115231424-8e868ca12c0f

replace github.com/golang/lint => github.com/golang/lint v0.0.0-20190227174305-8f45f776aaf1

replace github.com/nats-io/nats.go v1.8.2-0.20190607221125-9f4d16fe7c2d => github.com/nats-io/nats.go v1.8.1

require (
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/chzyer/logex v1.1.10 // indirect
	github.com/chzyer/test v0.0.0-20180213035817-a1ea475d72b1 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-redis/redis v6.15.2+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	github.com/go-telegram-bot-api/telegram-bot-api v4.6.4+incompatible // indirect
	github.com/gogo/protobuf v1.2.1
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.3.2
	github.com/google/uuid v1.1.1
	github.com/gorilla/mux v1.7.3
	github.com/gorilla/sessions v1.1.3
	github.com/gorilla/websocket v1.4.0
	github.com/grpc-ecosystem/grpc-gateway v1.9.5
	github.com/lusis/go-slackbot v0.0.0-20180109053408-401027ccfef5 // indirect
	github.com/lusis/slack-test v0.0.0-20190426140909-c40012f20018 // indirect
	github.com/micro/cli v0.2.0
	github.com/micro/go-config v1.1.1-0.20190603113547-03fb75f2f1e4
	github.com/micro/go-log v0.1.0
	github.com/micro/go-micro v1.9.0
	github.com/micro/go-plugins v1.2.0
	github.com/micro/go-web v1.0.0
	github.com/micro/micro v1.8.0
	github.com/montanaflynn/stats v0.5.0
	github.com/opentracing/opentracing-go v1.1.0
	github.com/uber/jaeger-client-go v2.16.0+incompatible
	go.uber.org/zap v1.10.0
	google.golang.org/genproto v0.0.0-20190801165951-fa694d86fc64
	google.golang.org/grpc v1.22.1
)
