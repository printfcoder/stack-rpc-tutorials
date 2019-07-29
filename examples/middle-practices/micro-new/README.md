# 新建模板

Micro中有`micro new`命令可以快速生成Micro风格的服务模板。本目录没有代码，主要是想让大家手动生成亲眼看一下效果。

## 使用

使用`micro new`很简单，我们先看一下它的使用方法

```text
$ micro new -h

NAME:
   micro new - Create a service template

USAGE:
   micro new [command options] [arguments...]

OPTIONS:
   --namespace value  Namespace for the service e.g com.example (default: "go.micro")
   --type value       Type of service e.g api, fnc, srv, web (default: "srv")
   --fqdn value       FQDN of service e.g com.example.srv.service (defaults to namespace.type.alias)
   --alias value      Alias is the short name used as part of combined name if specified
   --plugin value     Specify plugins e.g --plugin=registry=etcd:broker=nats or use flag multiple times
   --gopath           Create the service in the gopath. Defaults to true.
```

当我们创建新服务时，有两点我们要确认

|配置指令|作用|默认值|说明|
|---|---|---|---|
|--namespace|服务命令空间 |go.micro||
|--type|服务类型|srv|目前支持4种服务类型，分别是api、fnc(function)、srv(service)、web。|

其它选项都是可配可不配，一般使用默认即可

|配置指令|作用|默认值|说明|
|---|---|---|---|
|--fqdn|服务定义域，API需要通过该域找到服务|默认是使用服务的命令空间加上类型再加上别名||
|--alias|指定别名|声明则必填|使用单词，不要带任何标点符号，名称对Micro路由机制影响很大|
|--plugin|使用哪些插件|声明则必填|需要自选插件时使用|
|--gopath|是否使用GOPATH作为代码路径|true||

### 默认方式创建新服务

```bash
micro new github.com/micro-in-cn/tutorials/examples/middle-practices/micro-new/default
```

可以看到指令参数只有生成服务代码的路径，路径最后一个单词就是服务项目名，所以，最后一个单词一定<span style="color:red">不要加任何符号！</span>：

```text
Creating service go.micro.srv.default in /Users/me/workspace/go/src/github.com/micro-in-cn/tutorials/examples/middle-practices/micro-new/default

.
├── main.go
├── plugin.go
├── handler
│   └── example.go
├── subscriber
│   └── example.go
├── proto/example
│   └── example.proto
├── Dockerfile
├── Makefile
└── README.md

## 下面是提示要安装protobuf和使用protobuf指令生成类文件，已经安装有了protobuf可以忽略，直接切到项目目录，再执行protoc指令
download protobuf for micro:

brew install protobuf
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u github.com/micro/protoc-gen-micro

compile the proto file example.proto:

## 切目录，生成文件
cd /Users/me/workspace/go/src/github.com/micro-in-cn/tutorials/examples/middle-practices/micro-new/default
protoc --proto_path=. --go_out=. --micro_out=. proto/example/example.proto
```

生成的代码中，命令空间前缀为默认的**go.micro**，默认的服务类型为**srv**，服务别名（alias）为**default**

```go
package main
// ...
func main() {
    // New Service
    service := micro.NewService(
        micro.Name("go.micro.srv.default"),
        micro.Version("latest"),
    )
    // ...
}
```

### 指定命名空间

现在我们演示使用``--namespace``flag指定自己的命名空间，大家可以根据自己的域名定义合适的空间前缀，我们用micro官网的域名**mu.micro**代替

```bash
micro new --namespace=mu.micro github.com/micro-in-cn/tutorials/examples/middle-practices/micro-new/namespace
```

生成成功的消息与上面一样，不赘述，我们重点看下服务main方法内容：

```go
package main

// ...

func main() {
    // New Service
    service := micro.NewService(
        micro.Name("mu.micro.srv.namespace"),
        micro.Version("latest"),
    )

    // ...
    micro.RegisterSubscriber("mu.micro.srv.namespace", service.Server(), new(subscriber.Example))

    // Register Function as Subscriber
    micro.RegisterSubscriber("mu.micro.srv.namespace", service.Server(), subscriber.Handler)

    // ...
}
```

可以看到服务前缀由默认的**go.micro**改成了**mu.micro**，其它结构内容一致

### 指定服务类型

上面两个例子生成的都是默认的服务类型**srv**，现在我们演示指定为**api**类型。

```bash
micro new --type=api github.com/micro-in-cn/tutorials/examples/middle-practices/micro-new/apiType
```

```go
package main

// ...

func main() {
    // New Service
	service := micro.NewService(
		micro.Name("go.micro.api.apiType"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		// create wrap for the ApiType srv client
		micro.WrapHandler(client.ApiTypeWrapper(service)),
	)

	// Register Handler
	apiType.RegisterApiTypeHandler(service.Server(), new(handler.ApiType))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
```

如生成的代码所示，服务名中的类型部分变成了**api**。

需要注意的是Handler目录下的接口类，它的接口文件目录需要我们手动生成后复制进来

```go
import (
	// ...
	apiType "path/to/service/proto/apiType"
)
```

安装grpc工具链及protoc，[参考](https://grpc.io/docs/quickstart/go/)

生成接口文件

```bash
protoc --proto_path=.:$GOPATH/src --go_out=. --micro_out=. proto/apiType/apiType.proto
```

然后我们就可以看到在**proto/apiType/** 目录下生成了我们需要的接口文件：

```text
└── proto
    └── apiType
        ├── apiType.micro.go
        ├── apiType.pb.go
        └── apiType.proto
``` 

现在可以复制proto/api的相对GOPATH目录，然后把`path/to/service/proto/apiType`换成`github.com/micro-in-cn/tutorials/examples/middle-practices/micro-new/apiType/proto/apiType`。

### 指定FQDN

有些时候我们想要个性化的定义域，那么我们就可以指定`--fqdn`参数来声明。所谓定义域，默认情况下，服务全名就是**命名空间+服务类型+服务名**的组合，一旦设定`--fqdn`，它的值会替换默认值的组合。

下面我们把服务命名改为**mu.micro.fqdn.more**

```bash
micro new --fqdn=mu.micro.fqdn.more github.com/micro-in-cn/tutorials/examples/middle-practices/micro-new/fqdn
```

```go
package main

// ...

func main() {
    // New Service
    service := micro.NewService(
        micro.Name("mu.micro.fqdn.more"),
        micro.Version("latest"),
    )

    // ..
    
    // Register Struct as Subscriber
    micro.RegisterSubscriber("mu.micro.fqdn.more", service.Server(), new(subscriber.Example))

    // Register Function as Subscriber
    micro.RegisterSubscriber("mu.micro.fqdn.more", service.Server(), subscriber.Handler)

    // ...
}
```

或许有朋友会问，那`--namespace`和`--fqdn`起使用会怎么样，有兴趣的同学可以试试，这属于比较偏门且不正经的情况，我们不考虑。

### 指定别名

前面有提到路径最后一个单词就是服务项目名，或叫服务名。有时候我们不想让目录这个单词变成我们的服务名，比如下面的**micro**，我们想换成**orcim**，那我们就可以传入`--alias`指令。

```bash
micro new --alias=orcim github.com/micro-in-cn/tutorials/examples/middle-practices/micro-new/micro
```

```go
package main

// ...

func main() {
    // New Service
    service := micro.NewService(
        micro.Name("go.micro.srv.orcim"),
        micro.Version("latest"),
    )

    // ...
    // Register Struct as Subscriber
    micro.RegisterSubscriber("go.micro.srv.orcim", service.Server(), new(subscriber.Example))

    // Register Function as Subscriber
    micro.RegisterSubscriber("go.micro.srv.orcim", service.Server(), subscriber.Handler)

    // ...
}
```

生成的代码中，服务名已经由**micro**变成我们设置的**orcim**。

### 指定plugin

我们演示使用**etcd**插件和**kafka**插件

```bash
micro new --plugin=registry=etcd:broker=kafka github.com/micro-in-cn/tutorials/examples/middle-practices/micro-new/plugin
```

```go
package main

import (
    _ "github.com/micro/go-plugins/registry/etcd"
    _ "github.com/micro/go-plugins/broker/kafka"
)
```

在刚才的指令中我们声明注册中心使用**etcd**、broker消息代理使用**kafka**，然后new模板就会生成**plugin.go**文件，里面包含上面的代码。

### 不使用GOPATH

Micro new目前不支持自定义module目录，且Golang 1.11版本的**modules**现在还不100%成熟，大家就先默默使用GOPATH。