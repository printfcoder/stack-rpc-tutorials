# 客户端重试

演示如何使用客户端容错重试

- [client.go](client.go) 客户端
- [failureServer.go](failureServer.go) 错误服务端
- [successServer.go](successServer.go) 成功服务端

## 运行

运行failureServer.go 该服务端总是返回错误

```bash
go run failureServer.go
```

打开新窗口，运行successServer.go 该服务端返回成功

```bash
go run successServer.go
```

打开新窗口，运行客户端

```bash
go run client.go
```

见客户端与各服务端打印的日志
