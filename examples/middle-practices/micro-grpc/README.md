# GRPC

本篇给大家演示如何使用gRPC构建服务及使用gRPC gateway，代码请见[go-grpc](https://gitub.com/micro/go-grpc)

gRPC网关[gRPC-gateway](https://github.com/grpc-ecosystem/grpc-gateway)是[protoc](http://github.com/google/protobuf)的一个插件。

它遵循[gRPC](http://github.com/grpc/grpc-common)中对服务的定义，生成反向代理服务，这个代理就会把RESTful风格的JSON API转成gRPC请求。


## 预置条件

我们需要安装protobuf工具以及插件protoc-gen-grpc-gateway和protoc-gen-go。如果已经安装，请忽略以下步骤

```bash
mkdir tmp
cd tmp
git clone https://github.com/google/protobuf
cd protobuf
./autogen.sh
./configure
make
make check
sudo make install
```

```bash
$ go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
$ go get -u github.com/micro/protobuf/protoc-gen-go
```

## 内容

- [gateway](gateway) - gRPC网关
- [greeter](greeter) - 示例程序
  -- [greeter](greeter) - 示例程序
  -- [micro-cli](greeter/cli) - micro风格客户端
  -- [grpc-cli](greeter/grpc-cli) - grpc风格客户端
  -- [srv](greeter/srv) - 服务

## 开始上手

我们先创建需要的目录

```text
├── gateway    # 网关代码
├── greeter    # 示例程序
│   ├── cli    # 示例程序 服务端
│   └── srv    # 示例程序 客户端
└── proto      # proto 原型文件目录
    └── hello  # hello 原型文件及存根类目录
```

定义[hello.proto](proto/pb/greeter/greeter.proto)原型

```proto
syntax = "proto3";

package go.micro.srv.greeter;

service Say {
	rpc Hello(Request) returns (Response) {}
}

message Request {
	string name = 1;
}

message Response {
	string msg = 1;
}
```

定义好后，我们切到其上级目录，方便生成go文件：

```bash
# 生成纯grpc风格的go文件
protoc --proto_path=. greeter.proto --go_out=plugins=grpc:../../go/pure-grpc/
# 生成micro风格的go文件
protoc --proto_path=. --go_out=paths=source_relative:../../go/micro greeter.proto  --micro_out=paths=source_relative:../../go/micro
```

### srv服务

[srv](greeter/srv/main.go)代码

```
package main

import (
	"context"
	"log"
	"time"

	pb "github.com/micro-in-cn/tutorials/examples/middle-practices/micro-grpc/proto/go/micro"
	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
)

type Say struct{}

func (s *Say) Hello(ctx context.Context, req *hello.Request, rsp *hello.Response) error {
	log.Print("Received Say.Hello request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

func main() {
	service := grpc.NewService(
		micro.Name("go.micro.srv.greeter"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.Address(":9090"),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	pb.RegisterSayHandler(service.Server(), new(Say))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
```

### micro风格客户端

[micro-cli](greeter/cli/main.go)

```go
package main

import (
	"context"
	"fmt"

	pb "github.com/micro-in-cn/tutorials/examples/middle-practices/micro-grpc/proto/go/micro"
	"github.com/micro/go-grpc"
	"github.com/micro/go-micro/metadata"
)

func main() {
	service := grpc.NewService()
	service.Init()

	// use the generated client stub
	cl := pb.NewSayService("go.micro.srv.greeter", service.Client())

	// Set arbitrary headers in context
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id": "john",
		"X-From-Id": "script",
	})

	rsp, err := cl.Hello(ctx, &pb.Request{
		Name: "John",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Msg)
}
```

### grpc风格客户端

[grpc-client](greeter/grpc-cli/main.go)

```go
const (
	address     = "localhost:9090"
	defaultName = "我是来自grpc风格的客户端请求！"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSayClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Hello(ctx, &pb.Request{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Logf("Greeting: %s", r.Msg)
}
```

### 运行程序

#### srv

我们切到srv目录，并运行服务

```bash
cd greeter/srv
go run main.go
```

程序启动并打印如下日志，我们记下grpc服务端口9090

```bash
2019/07/15 23:20:45 Server [grpc] Listening on [::]:9090
2019/07/15 23:20:45 Broker Listening on [::]:52507
2019/07/15 23:20:45 Broker [grpc] Listening on [::]:52507
2019/07/15 23:20:45 Registering node: go.micro.srv.greeter-13c29b5f-3f91-4f09-a801-f2fdb1710f2d
```

#### micro

切到micro风格客户端目录，并运行

```bash
cd greeter/cli
go run main.go
```

尔后程序会打印如下信息：

```bash
Hello 我是来自micro风格的客户端请求
```

#### grpc

在运行之前，我们要去服务端地址告诉客户端，比如我们刚启动的srv的grpc端口是52506

```go
const (
	address     = "localhost:52506"
	defaultName = "我是来自grpc风格的客户端请求！"
)
```

随后切到grpc风格客户端目录，并运行

```bash
cd greeter/grpc-cli
go run main.go
```

程序会打印如下信息：

```bash
2019/07/15 23:24:08 Greeting: Hello 我是来自grpc风格的客户端请求！
```


## grpc Gateway

如不了解grpc网关，请先点击[grpc网关](https://github.com/grpc-ecosystem/grpc-gateway)。

grpc网关其实就是把grpc服务转成普通的http服务的反向代理接入层，通过它把grpc服务的rpc类型的接口暴露成restful风格的路由。

我们先生成gateway的接口声明代码：

```bash
# 生成grpc接口声明
cd proto/pb/gateway
protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:../../go/grpc-gw \
  greeter/greeter.proto

# 生成gateway声明
protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --grpc-gateway_out=logtostderr=true:../../go/grpc-gw  \
  greeter/greeter.proto
```

### 调用

```bash
curl -d '{"name": "john"}' http://localhost:8080/greeter/hello
```

控制台便会打印

```bash
{"msg":"Hello john"}
```
