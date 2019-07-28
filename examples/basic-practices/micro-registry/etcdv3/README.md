# etcdv3

本篇演示如何使用etcdv3。

## 内容

- server.go - 服务端
- client.go - 客户端
- plugins.go - etcdv3插件

## 运行

运行服务端

```shell
go run server.go plugins.go --registry=etcdv3
```

运行客户端

```shell
go run client.go plugins.go --registry=etcdv3
```
