# WRAP CALL

**WrapCall**负责在客户端调用服务前进行包装处理，所有在WrapCall中的包装器都会被执行。

## 开始

运行**server**服务端：

```
go run server.go
```

切换终端窗口：

```
go run client.go
```

服务端窗口会打印我们在WrapCall中声明的头数据：

```
INFO[0020] [Hello] call-wrapped1: call-wrapped-value1    source="server.go:19"
INFO[0020] [Hello] call-wrapped2: call-wrapped-value2    source="server.go:20"
```

任何发起请求前的处理，都可以在Call中包装。