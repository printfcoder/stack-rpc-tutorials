# NSQ Pubsub

本篇演示如何使用NSQ消息中间件来编写Micro的Pubsub服务，本篇由[Jayn Yang](https://github.com/jayn1985)提供，略作结构上的删改。

## 预置条件：

安装NSQ，参考官网[安装](https://nsq.io/deployment/installing.html)，或者使用[Docker](https://nsq.io/deployment/docker.html)

我们假设读者的NSQ工作在本地标准地址：

- NSQ：127.0.0.1:4150

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
2019/08/05 23:54:48 Receive info: Id 0 & Timestamp 1565020484
2019/08/05 23:54:48 Receive info: Id 1 & Timestamp 1565020485
2019/08/05 23:54:48 Receive info: Id 2 & Timestamp 1565020486
2019/08/05 23:54:48 Receive info: Id 3 & Timestamp 1565020487
2019/08/05 23:54:48 Receive info: Id 4 & Timestamp 1565020488
2019/08/05 23:54:49 Receive info: Id 5 & Timestamp 1565020489
2019/08/05 23:54:50 Receive info: Id 6 & Timestamp 1565020490
```