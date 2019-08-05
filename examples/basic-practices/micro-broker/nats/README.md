# NATs Pubsub

本篇演示如何使用NATs消息中间件来编写Micro的Pubsub服务，原稿由[Bruce Wang](https://github.com/BruceWangNo1/go-micro-pubsub-with-nats)提供，略作删改。

## 预置条件

安装NatsServer，教程参考，[安装NATs Server](https://nats-io.github.io/docs/nats_server/installation.html)

然后运行NATs服务（操作系统不一运行指令不一样）。

```bash
nats-server
```

## 目录

- [cli](cli) 客户端
- [srv](srv) 服务端

## 运行示例服务

先打开一个终端窗口，切到我们的示例**服务**目录：

```bash
cd srv
go run main.go --broker=nats --broker_address=127.0.0.1:4222
```

再打开一个终端窗口，切到我们的示例**客户端**目录：

```bash
cd cli
go run main.go --broker=nats --broker_address=127.0.0.1:4222
```

然后我们可以在两个窗口中看到打印的信息：

cli 窗口

```text
2019/08/01 14:07:14 publishing {Id:9beb1a7e-b422-11e9-92ac-acde48001122 Timestamp:1564639634 Message:如果你看到了消息 go.micro.pubsub.topic.event, '那是因为我一直爱着你 XXX_NoUnkeyedLiteral:{} XXX_unrecognized:[] XXX_sizecache:0}
...
```

srv 窗口

```text
2019/08/01 14:07:09 [sub] 收到消息，请查收: map[id:3], {"id":"98f050d2-b422-11e9-92ac-acde48001122","timestamp":1564639629,"message":"如果你看到了消息 go.micro.pubsub.topic.event, '那是因为我一直爱着你"}
...
```

Thanks: https://github.com/BruceWangNo1/go-micro-pubsub-with-nats
