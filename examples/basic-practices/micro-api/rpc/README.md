# RPC API

本示例我们介绍如何使用RPC handler模式的**Micro API**，以下简称**API**。

该模式下允许我们通过RPC的方式把HTTP请求转发到go-micro微服务上。

需要提醒的是，RPC模式下**API**只接收POST方式的请求，并且只支付内容格式**content-type**为**application/json**或者**application/protobuf**。

## 使用方法

使用protoc生成go代码
```
protoc  --go_out=. --micro_out=. proto/api.proto
```

以rpc模式运行**API**

```
micro api --handler=rpc
```

```
go run rpc.go
```

当我们POST请求到 **/example/call**时，**API**会将它转成RPC转发到**go.micro.api.example**服务的**Example.Call**接口上。

```
curl -H 'Content-Type: application/json' -d '{"name": "小小先"}' "http://localhost:8080/example/call"
```

同样，POST请求到 **/example/foo/bar**时，**API**会将它转成RPC转发到**go.micro.api.example**服务的**Foo.Bar**接口上。

```
curl -H 'Content-Type: application/json' -d '{}' http://localhost:8080/example/foo/bar
```
