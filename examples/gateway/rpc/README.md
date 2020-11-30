# RPC API

使用Stack Gateway插件暴露RPC服务

### 运行RPC，参考[RPC-Demo](../../service/rpc)

```bash
go run server.go
```

### 运行Gateway

运行本示例中的main.go

```bash
go run main.go --gateway_handler=rpc --gateway_namespace=stack.rpc
```

### 请求

```bash
curl POST '127.0.0.1:8080/greeter/hello' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Curl Client"
}'
```