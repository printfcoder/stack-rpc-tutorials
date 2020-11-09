# RPC API

本示例我们介绍如何使用Stack-RPC来编写RPC服务。

示例中我们有两个程序：

- [client](./client.go) 客户端，负责请求server
- [server](./server.go) 服务端，负责提供Greeter服务

## 使用方法

使用protoc生成go代码，proto文件在[proto](../../proto/service/rpc/)目录，内容如下：

```proto
syntax = "proto3";

service Greeter {
    rpc Hello(HelloRequest) returns (HelloResponse) {}
}

message HelloRequest {
    string name = 1;
}

message HelloResponse {
    string greeting = 1;
}
```

它提供一个服务叫Greeter，方法叫Hello，入参和出参分别是HelloRequest与HelloResponse

生成proto文件，没有protoc及stack插件，可以参考[准备安装]()

```
protoc  --go_out=. --stack_out=. proto/rpc.proto
```

以rpc模式运行**API**

```
micro api --handler=rpc --namespace=go.micro.rpc
```

```
go run rpc.go
```

当我们POST请求到 **/example/call**时，**API**会将它转成RPC转发到**go.micro.rpc.example**服务的**Example.Call**接口上。

```
curl -H 'Content-Type: application/json' -d '{"name": "小小先"}' "http://localhost:8080/example/call"
```

同样，POST请求到 **/example/foo/bar**时，**API**会将它转成RPC转发到**go.micro.api.example**服务的**Foo.Bar**接口上。

```
curl -H 'Content-Type: application/json' -d '{}' http://localhost:8080/example/foo/bar
```
