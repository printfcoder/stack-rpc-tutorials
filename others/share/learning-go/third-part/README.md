# 夜读第三期

## MICRO API

启动API：

```shell
micro api --handler=rpc --namespace=go.micro.api
```

启动rpc服务

```shell
go run rpc/main.go
```

启动rpc服务v2

```shell
go run rpc-v2/main.go
```

调用Greeter

```shell
curl http://127.0.0.1:8080/learning/greeter/hi  --header 'Content-Type: application/x-www-form-urlencoded' --data '{"name":"Micro"}'
```

调用Learning

```shell
curl http://127.0.0.1:8080/learning/hi  --header 'Content-Type: application/x-www-form-urlencoded' --data '{"name":"Micro"}'
```

调用Learning V2

```shell
curl http://127.0.0.1:8080/v2/learning/hi  --header 'Content-Type: application/x-www-form-urlencoded' --data '{"name":"Micro"}'
```

## Micro CLI

### micro list services

```bash
micro list services
```

```bash
micro --registry=etcd list services
```

### micro get/health service

```bash
micro get service go.micro.api.learning
```

### micro call 

```bash
micro call go.micro.api.learning Learning.Hi '{"name":"Micro"}'
```

### micro register

**mdns doest support `micro register`**

micro api 注册到etcd

```shell
micro --registry=etcd api --handler=rpc --namespace=go.micro.api
```

v2 注册到mdns

复制上面的节点id与端口，

```bash
micro register service '{"name": "go.micro.api.v2.learning", "version": "v2", "nodes": [{"id": "83f0aeb8-0c38-40bb-9aa2-f35a472cc4b7", "address": "127.0.0.1:54235"}]}'
```

RPC调用Greeter

```bash
micro call go.micro.api.v2.learning Learning.Hi '{"name":"Micro"}'
```