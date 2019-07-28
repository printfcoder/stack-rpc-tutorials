# Web API

本篇演示如何使用Web模式下的**Micro API**，以下简称**API**。

在web代理模式下运行**API**，我们可以将**API**作为反向代理提供下游服务的http接口，示例中我们演示使用websocket。

**API**会向注册中心查询服务信息，将请求路由转向合适的后台服务上。故而我们直接使用go-web作为后台服务，因为它可以直接注册，为了方便我们不直接从头写可以注册的服务。

## 使用方法

以web模式运行**API**，因为我们的应用是在web空间下，所以我们把api的启动空间设置为**go.micro.web**

```
micro api --handler=web --namespace=go.micro.web
```

运行web应用

```
go run web.go
```

## 演示

打开 http://127.0.0.1:8080/websocket/

默认会打开websocket连接。在Name栏中输入文本，点击send按钮便可以与后台websocket服务交互。