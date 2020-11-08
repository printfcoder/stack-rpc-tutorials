# Event

本篇演示如何使用Event模式下的**Micro API**，以下简称**API**。

**API**会把http请求映射到匹配的事件处理服务上

## 运行

因为我们在代码中声明了事件主题topic是`go.micro.evt.user`，即是说事件服务的命名所属空间是`go.micro.evt`，所以我们的**API**也要是这个命名空间，这样**API**才能找到它。

```bash
micro api --handler=event --namespace=go.micro.evt
```

运行服务

```
go run main.go
```

发送事件

```bash
curl -d '{"message": "Hello, Micro中国"}' http://localhost:8080/user/login
```
