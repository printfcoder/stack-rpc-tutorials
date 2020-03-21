# micro 跨域

V2.2.0及以前的版本，实现跨域，只需要在 micro api（micro API 网关）注册相关插件，并且通过命令行指定跨域相关信息即可。

```go
plugin.Register(cors.NewPlugin())
```

## 运行API网关（V2.2.0及以前的版本）

默认的API是没有跨域的，我们要加上插件，详见代码[api](./api/main.go)

```bash
go run main.go \
--cors-allowed-headers="Content-Type,X-Token" \
--cors-allowed-origins="*" \
--cors-allowed-methods="OPTIONS,DELETE,GET,POST" api --handler=api
```

## 运行API网关（V2.3.0及以上的版本）

```bash
micro api --enable_cors=true --handler=api
```

## 运行API服务

API服务只是为了有个URL能让我们请求

```bash
go run ../micro-api/api/api.go
```

## 运行Web

```bash
go run web/main.go
```

打开浏览器，输入http://localhost:9090，点击按钮即可向端口8080跨域请求数据