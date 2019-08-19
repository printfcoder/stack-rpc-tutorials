# 客户端重试

演示如何使用客户端容错重试

[客户端](./client.go) 
[错误服务端](./FailureServer.go) 
[成功服务端](./SuccessServer.go) 

## 运行

运行FailureServer.go 该服务端总是返回错误

```bash
go run FailureServer.go
```

打开新窗口，运行SuccessServer.go 该服务端返回成功

```bash
go run FailureServer.go
```

打开新窗口，运行客户端

```bash
go run client.go
```

见客户端与各服务端打印的日志
