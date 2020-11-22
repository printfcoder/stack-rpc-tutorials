# RPC API

本示例我们介绍如何使用Stack-RPC来编写RPC服务。

示例中我们有两个程序：

- [client](./client.go) 客户端，负责请求server
- [server](./server.go) 服务端，负责提供Greeter服务

## 使用方法

使用protoc生成go代码，proto文件在[proto](../../proto/service/rpc)目录，内容如下：

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

生成proto文件（已生成好，但大家可自行尝试），没有protoc及stack插件，可以参考[准备安装](http://microhq.cn/docs/stack-rpc/prepare-env-cn)

切到proto的rpc目录，再执行：

```
protoc  --go_out=. --stack_out=. rpc/rpc.proto
```

## 编写Server

```golang
// 服务类
type Greeter struct {
}

// 实现proto中的Hello接口
func (g Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello! " + req.Name
	return nil
}

func main() {
    // 实例化服务，并命名为stack.rpc.greeter
	service := stack.NewService(
		stack.Name("stack.rpc.greeter"),
	)
    // 初始化服务
	service.Init()

	// 将Greeter注册到服务上
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

    // 运行服务
	if err := service.Run(); err != nil {
		logger.Error(err)
	}
}
```

运行server:

```
go run server.go
```

打开另一窗口，运行client:

```
go run client.go
```
