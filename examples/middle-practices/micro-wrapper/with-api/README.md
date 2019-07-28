## API

我们在这里演示通过`micro api`调用有wrapper包装的接口

## 使用方式

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