# Timeout

本篇演示如何指定Timeout

## 内容

- main.go - 服务端
- client.go - 客户端

micro支持以下方式使用Timeout

环境变量：

```bash
export MICRO_CLIENT_REQUEST_TIMEOUT=10s
```

参数：

```golang
    cli := client.NewClient(
		client.RequestTimeout(time.Second * 15),
	)

	service := micro.NewService(
		micro.Name("timeout.client"),
		micro.Client(cli))
```

命令行：

```bash
go run client.go --client_request_timeout=10s
```

运行服务端

```shell
go run server.go
```

运行客户端（参数方式）

```shell
go run client.go
```

运行客户端（命令行方式）

删除掉客户端部分代码：

```go
func main() {
	// 定义服务，可以传入其它可选参数
	service := micro.NewService(
		micro.Name("timeout.client"),
	)
	service.Init()

	// 创建客户端
	greeter := proto.NewGreeterService("timeout.service", service.Client())

	// 调用greeter服务
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "Micro中国"})
	if err != nil {
		fmt.Println(err)
		return
	}

	// 打印响应结果
	fmt.Println(rsp.Greeting)
}
```

```shell
go run client.go --client_request_timeout=5s
```

运行客户端（环境变量）

```bash
export MICRO_CLIENT_REQUEST_TIMEOUT=10s
go run client.go
```