# Proxy API

本篇演示如何使用proxy模式下的**Micro API**，以下简称**API**。

在proxy代理模式下运行**API**，我们可以自行决定使用何种语言或类库来写我们的接口层应用。

**API**会向注册中心查询服务信息，将请求路由转向合适的后台服务上。故而我们直接使用go-web作为后台服务，因为它可以直接注册，为了方便我们不直接从头写可以注册的服务。

## 使用方法

以http模式运行**API**

```
micro api --handler=http
```

运行代理应用

```
go run proxy.go
```

### 示例一 /example/call

发送请求到 **/example/call**，该请求会被**API**反向代理到**go.micro.api.example**服务的 **/example/call**路由

```
curl "http://localhost:8080/example/call?name=micro"
```

### 示例二 /example/foo/bar

POST请求到 **/example/foo/bar**会调用**go.micro.api.example**的 **/example/foo/bar**路由。

```
 curl -H 'Content-Type: application/json' -d '{"name": "micro"}' http://localhost:8080/example/foo/bar
```

### 示例三 文件上传 /example/foo/upload

我们可以请求`http://localhost:8080/example/foo/upload`，获取上传页面，选择适当的文件上传，测试上传功能。为了方便和直观，请保存上传保存的目录存在，且上传小文件