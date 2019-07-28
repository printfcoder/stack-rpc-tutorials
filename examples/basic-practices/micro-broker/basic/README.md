# Broker

Micro中的[broker](https://godoc.org/github.com/micro/go-micro/broker#Broker)用于代理消息的发布与订阅。

本章节只讲如何使用broker的基础方式

## 内容

- main.go - mian程序运行两个协程10秒钟，各自负责发布与订阅消息。

## 运行程序

```bash
go run main.go
```