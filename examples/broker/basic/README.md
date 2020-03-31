# Broker

Micro中的[broker](https://godoc.org/github.com/micro/go-micro/broker#Broker)用于代理消息的发布与订阅。

## 内容

- main.go - main程序运行两个协程20秒钟，各自负责发布与订阅消息。

## 运行程序

如果使用默认的http broker，请运行：

```bash
go run main.go
```

如果想使用其他消息队列服务，例如nats，请运行：

```bash
export MICRO_BROKER=nats
go run main.go
```

或者：

```bash
go run main.go --broker=nats
```

或者：

```bash
go run main.go --broker=nats --broker_address=127.0.0.1:4222
```
