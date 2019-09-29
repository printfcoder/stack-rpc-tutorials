# NSQ Pubsub

本篇演示如何使用Kafka消息中间件来编写Micro的Pubsub服务，本篇由[Jayn Yang](https://github.com/jayn1985)提供，略作结构上的删改。

## 预置条件：

安装 Kafka

## 目录

- [cli](cli) 客户端
- [srv](srv) 服务端

## 运行程序

### 客户端

```bash
cd cli
go run client.go
```

### 服务端

```bash
cd srv
go run server.go
```

### 日志

我们可以看到日志打印如下信息

```text
2019/09/29 22:30:13 Receive info: Id 0 & Timestamp 1569767413
2019/09/29 22:30:14 Receive info: Id 1 & Timestamp 1569767414
2019/09/29 22:30:15 Receive info: Id 2 & Timestamp 1569767415
2019/09/29 22:30:16 Receive info: Id 3 & Timestamp 1569767416
2019/09/29 22:30:17 Receive info: Id 4 & Timestamp 1569767417
2019/09/29 22:30:18 Receive info: Id 5 & Timestamp 1569767418
2019/09/29 22:30:19 Receive info: Id 6 & Timestamp 1569767419
2019/09/29 22:30:20 Receive info: Id 7 & Timestamp 1569767420
```