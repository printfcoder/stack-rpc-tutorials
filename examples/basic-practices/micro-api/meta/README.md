# Meta API

本篇演示如何使用metadata模式下的**Micro API**，以下简称**API**。go-micro支持将请求路由到服务元数据声明的方法，也即是基于元数据的服务发现。

## 使用方法
使用protoc生成go代码
```
protoc --go_out=. --micro_out=. proto/api.proto
```

运行API网关，可以看到，**API**启动时，并没有声明handler模式，故而使用的**RPC**模式。所以**Meta API**其实是在RPC模式的基础上，通过在接口层声明端点元数据而指定服务的。

```
micro api
```

运行示例程序，在代码中注册服务时，我们在endpoint参数中写入了元数据，声明接口为 **/example**和 **/foo/bar**

```
go run meta.go
```

向 **/example** POST请求时，会被转到**go.micro.api.example**的**Example.Call**方法。

```
curl -H 'Content-Type: application/json' -d '{"name": "john"}' "http://localhost:8080/example"
```

向 **/example** POST请求时，会被转到**go.micro.api.example**的**Foo.Bar**方法。

```
curl -H 'Content-Type: application/json' -d '{}' http://localhost:8080/foo/bar
```
