# micro plugins

## 跨域

实现跨域，只需要在 micro api（micro API 网关）注册相关插件，并且通过命令行指定跨域相关信息即可。

```go
plugin.Register(cors.NewPlugin())
```

```bash
go run main.go \
--cors-allowed-headers="Content-Type,X-Token" \
--cors-allowed-origins="*" \
--cors-allowed-methods="OPTIONS,DELETE,GET,POST" api
```
