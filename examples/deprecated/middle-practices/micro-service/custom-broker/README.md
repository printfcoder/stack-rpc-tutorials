# 自定义Broker

本篇演示如何自定义Broker。

## 内容

- server.go - 服务端
- pub.go - 发布端

**pub.go**发布端负责发布两个主题消息**go.micro.topic.custom1**和**go.micro.topic.custom2**

**server.go**发布端负责订阅两个主题消息**go.micro.topic.custom1**和**go.micro.topic.custom2**

## 运行

运行服务端

```shell
go run server.go
```

运行发布端

```shell
go run pub.go
```
