# 第一章 用户服务

本章节我们实现用户服务，用户服务分为两层，web层（user-web）与服务层（user-service），前者提供http接口，后者向web提供RPC服务。

- user-web 以下简称web
- user-service 以下简称service

web服务主要向用户提供如下接口

- 登录与token颁发
- 鉴权

我们不提供注册接口，一来增加不必要的代码量，我们的核心还是介绍如何使用Micro组件。

## 开始编写

我们先从下往上编写，也就是从服务层**user-service**开始

### user-service

Micro有提供代码生成器指令[**new**][micro-new]，它可以新建服务模板代码，下面我们使用它来创建用户服务

需要有个说明的地方，因为我们现在执行的命令是直接将代码生成到当前的目录，所以大家在自己运行时，请修改到自己的路径，以免覆盖示例程序。

好，我们开始

#### 新建模板

```bash
micro new --namespace=mu.micro.book.user --type=srv --alias=service github.com/micro-in-cn/micro-tutorials/microservice-in-micro/part1/user-service
```

模板生成在**user-service**目录，其结构如下

```text
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

```

生成的部分代码我们用不到，把它们删掉，加上我们需要的文件。结构如下

```text
.
├── main.go
├── plugin.go
├── handler
│   └── service.go
├── proto/service
│   └── service.proto
├── Dockerfile
├── Makefile
└── README.md

```

#### 定义User原型

我们需要在service.proto中定义User原型，暂且定义以下字段，足够登录与显示用户基本信息即可;

```proto
syntax = "proto3";

package mu.micro.book.user.srv.service;

service Service {
    rpc QueryUser (Request) returns (Response) {
    }
}

message User {
    int64 id = 1;
    string name = 2;
    string pwd = 3;
}

message Request {
    string userID = 1;
    string userName = 2;
    string userPwd = 3;
}

message Response {
    User user = 1;
}
```

上面我们定义了User服务的基本原型结构，包含用户**User**，请求**Request**与响应结构**Response**，还定义了查询用户的方法**QueryUser**。

下面我们生成类型与服务方法：

```bash
protoc --proto_path=. --go_out=. --micro_out=. proto/service/service.proto
```

### user-web

## 延伸阅读

[使用Micro模板新建服务][micro-new]

[micro-new]: https://github.com/micro-in-cn/micro-all-in-one/tree/master/middle-practices/micro-new