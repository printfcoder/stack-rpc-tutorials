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

Micro有提供代码生成器指令[**new**][micro-new]，它可以新建服务模板代码，把基本所需的目录结构建好，省去大家挖坑的时间。

下面我们使用它来创建用户服务。

- <span style="color:red">*</span>需要有个说明的地方，因为我们现在执行的命令是直接将代码生成到当前的目录，所以大家在自己运行时，请修改到自己的路径，以免覆盖示例程序，引起不必要的麻烦。

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

#### 定义User原型

我们需要在service.proto中定义User原型，暂且定义以下字段，足够登录，显示用户基本信息、异常信息即可;

```proto
syntax = "proto3";

package mu.micro.book.user.srv.service;

service Service {
    rpc QueryUserByName (Request) returns (Response) {
    }
}

message User {
    int64 id = 1;
    string name = 2;
    string pwd = 3;
    uint64 createdTime = 4;
    uint64 updatedTime = 5;
}

message Error {
    int32 code = 1;
    string detail = 2;
}

message Request {
    string userID = 1;
    string userName = 2;
    string userPwd = 3;
}

message Response {
    bool success = 1;
    Error error = 2;
    User user = 3;
}
```

上面我们定义了User服务的基本原型结构，包含用户**User**，请求**Request**与响应结构**Response**，还定义了查询用户的方法**QueryUser**。

下面我们生成类型与服务方法：

```bash
protoc --proto_path=. --go_out=. --micro_out=. proto/service/service.proto
```

```go
package main

// ...

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("mu.micro.book.user.srv.service"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	s.RegisterServiceHandler(service.Server(), new(handler.Service))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("mu.micro.book.user.srv.service", service.Server(), new(subscriber.Service))

	// Register Function as Subscriber
	micro.RegisterSubscriber("mu.micro.book.user.srv.service", service.Server(), subscriber.Service)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

```

生成的**main**方法比较简单，根据我们当前的需要，我们把不要的pubsub（发布订阅）都删掉，变成：

```go
package main

// ...

func main() {
	// New Service   新建服务
	service := micro.NewService(
		micro.Name("mu.micro.book.user.srv.service"),
		micro.Version("latest"),
	)

	// Initialise service  初始化服务
	service.Init()

	// Register Handler   注册服务
	s.RegisterServiceHandler(service.Server(), new(handler.Service))

	// Run service    启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

```

朋友们可能已经发现，如果我们要从数据库里获取数据，模块提供的代码是远远不够的，那下面我们就开始真正开始编写代码。

#### 开始写代码

上面的生成的目录部分，我们需要手动改一下：

生成的部分代码我们用不到，把它们删掉，加上我们需要的文件与目录。结构如下

```text
.
├── main.go
├── plugin.go
├── basic
│   └── config               * 配置类
│   │   └── config.go        * 初始化配置类
│   │   └── config_consul.go * consul配置结构
│   │   └── config_mysql.go  * mysql配置
│   │   └── profiles.go      * 配置文件树辅助类
│   └── db                   * 数据库相关
│   │    └── db.go           * 初始化数据库
│   │    └── mysql.go        * mysql数据库相关
│   └── basic                * 初始化基础组件
├── config                   * 配置文件目录
├── handler
│   └── service.go           * 将名称改为service
├── model                    * 增加模型层，用于与数据库交换数据
│   └── user                 * 用户模型类
│   │   └── user.go          * 初始化用户模型类
│   │   └── user_get.go      * 封装获取用户数据类业务
│   └── model.go             * 初始化模型层
├── proto/service    
│   └── service.proto        * 将名称改为service
├── Dockerfile
├── Makefile
└── README.md
```

其中加`*`的便是我们修改过的结构，其后跟的描述是目录或文件的功能或作用。可能大家会觉得改动这么大，模板命令还有什么用呢？

其实模板只是生成基础目录，把大家引进一个风格的项目中，这样管理起来会轻松许多。下面我们解释一下为什么要新增两个目录：**basic**，**model**和**config**。

**basic**和**model**其实和Micro无关，只是为了满足我们为**user-service**的业务定位，它是一个**MVC**应用，其中C交给了**user-web**，剩下的**MV**才是它的主要功能。

- **basic** 负责初始化基础组件，比如数据库、配置等

- **model** 负责封装业务逻辑

- **config** 配置文件目录，现在我们还没用配置中心，暂先用文件的方式

有朋友会问，那**handler**目录呢？**user-service**本质上也是一个MVC应用，它弱化了C成handler，只负责接收请求，不改动业务数据**值**，但可能改动结构以便回传。

下面我们开始处理业务方面的东西

##### 创建User表

我们选用Mysql作为数据库，以下是建表语句，完整sql可以在[文档](./docs/schema.sql)目录找到：

```sql
CREATE TABLE `user`
(
    `id`           int(10) unsigned                                              NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id`      int(10) unsigned                                                       DEFAULT NULL COMMENT '用户id',
    `user_name`    varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci  NOT NULL COMMENT '用户名',
    `pwd`          varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '密码',
    `created_time` timestamp(3)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_time` timestamp(3)                                                  NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
    PRIMARY KEY (`id`),
    UNIQUE KEY `user_user_name_uindex` (`user_name`),
    UNIQUE KEY `user_user_id_uindex` (`user_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='用户表';
```

预置一条数据，<span style="color:red">为了简化，我们的账户密码暂时使用明文</span>，后面的章节会加盐hash后再存储、匹配。

```sql
INSERT INTO user (user_id, user_name, pwd) VALUE (10001, 'micro', '123');
```

##### 配置

配置目录

### user-web

## 延伸阅读

[使用Micro模板新建服务][micro-new]

[micro-new]: https://github.com/micro-in-cn/micro-all-in-one/tree/master/middle-practices/micro-new