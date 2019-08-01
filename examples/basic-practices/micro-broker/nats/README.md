# NATs Pubsub

本篇演示如何使用NATs使用消息中间件来编写Micro的Pubsub服务，本篇由[Bruce Wang](https://github.com/BruceWangNo1/go-micro-pubsub-with-nats)提供，略作删改。

- [cli](cli) 客户端
- [srv](srv) 服务端## Generating srv Service Template


## 预置条件

安装NatsServer，教程参考，[安装NATs Server](https://nats-io.github.io/docs/nats_server/installation.html)

然后运行NATs服务（操作系统不一运行指令不一样）。

```bash
nats-server
```

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

Thanks: https://github.com/BruceWangNo1/go-micro-pubsub-with-nats