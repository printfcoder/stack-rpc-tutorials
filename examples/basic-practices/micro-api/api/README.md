## API

本示例介绍使用**Micro API**（以下简称**API**）中的请求处理类型**api**，这里后面的api意思不是Micro API工具，而是指具体的对外的API接口服务。

另外，我们专门定义有api请求响应的proto文件，[api.Request/Response](https://github.com/micro/go-micro/blob/master/api/proto/api.proto)，

要使用**api**类型的**API**服务，我们得使用这些proto原型。

## 使用方式

使用protoc生成相应的代码
```
protoc --proto_path=import_proto_path:. --go_out=. --micro_out=. proto/api.proto
```
>  会在 import_proto_path 这个路径下搜索proto/api.proto中导入proto文件

运行**API**网关，我们传入api指令运行：

```
micro api --handler=api
```

再运行本api服务

```
go run api.go
```

## 调用服务

通过URL **/example/call**，就会调用**go.micro.api.example**服务的**Example.Call**接口

请求头的数据会被传到最终调用的接口

```
curl -H 'head-1: I am a header' "http://localhost:8080/example/call?name=john"
```

而POST路由 **/example/foo/bar**，可以通过它调用**go.micro.api.example**服务的**Foo.Bar**接口

```
curl -H 'Content-Type: application/json' -d '{data:123}' http://localhost:8080/example/foo/bar
```

## 设置命名空间

可以通过`--namespace`指定服务命令空间

```
micro api --handler=api --namespace=com.foobar.api
```

或者通过环境变量的方式

```
MICRO_API_NAMESPACE=com.foobar.api micro api --handler=api
```

切记，如果启动时指定命名空间，则代码中的服务名也要注意同步修改前缀，即把**micro.Name**的参数改成对应的命令空间前缀，以便**API**通过解析路由找到它。

```
service := micro.NewService(
        micro.Name("com.foobar.api.example"),
)
```   
