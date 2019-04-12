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

##### 基础组件

[基础组件](./user-service/basic)目前主要的功能是初始化配置与数据库。它的入口代码是一个**Init**初始化方法，负责初始化其下所有组件。

```go
package basic

import (
	"github.com/micro-in-cn/micro-tutorials/microservice-in-micro/part1/user-service/basic/config"
	"github.com/micro-in-cn/micro-tutorials/microservice-in-micro/part1/user-service/basic/db"
)

func Init() {
	config.InitConfig()
	db.InitDB()
}
```

###### 配置

初始化配置的过程大致如下：

|顺序|过程|说明|
|---|---|---|
|1|加载application.yml|读取conf目录下application.yml文件|
|2|解析profiles属性|如果有该属性则找到include值，该值就是指定需要引入的conf下的配置文件|
|3|解析include|解析出include配置【值】，并组合成文件名，文件名规则为[application-值.yml]|
|4|读取include声明文件|读取配置文件值|
|5|解析配置|将配置文件中的值解析到配置对象中|

下面是它的核心代码

```go
// InitConfig 初始化配置
func InitConfig() {

	m.Lock()
	defer m.Unlock()

	if inited {
		log.Fatal(fmt.Errorf("[InitConfig] 配置已经初始化过"))
		return
	}

	// 加载yml配置
	// 先加载基础配置
	appPath, _ := filepath.Abs(filepath.Dir(filepath.Join("./", string(filepath.Separator))))

	pt := filepath.Join(appPath, "conf")
	os.Chdir(appPath)

	// 找到application.yml文件
	if err = config.Load(file.NewSource(file.WithPath(pt + "/application.yml"))); err != nil {
		panic(err)
	}

	// 找到需要引入的新配置文件
	if err = config.Get(defaultRootPath, "profiles").Scan(&profiles); err != nil {
		panic(err)
	}

	log.Logf("[InitConfig] 加载配置文件：path: %s, %+v\n", pt+"/application.yml", profiles)

	// 开始导入新文件
	if len(profiles.GetInclude()) > 0 {
		include := strings.Split(profiles.GetInclude(), ",")

		sources := make([]source.Source, len(include))
		for i := 0; i < len(include); i++ {
			filePath := pt + string(filepath.Separator) + defaultConfigFilePrefix + strings.TrimSpace(include[i]) + ".yml"
			fmt.Printf(filePath + "\n")
			sources[i] = file.NewSource(file.WithPath(filePath))
		}

		// 加载include的文件
		if err = config.Load(sources...); err != nil {
			panic(err)
		}
	}

	// 赋值
	config.Get(defaultRootPath, "consul").Scan(&consulConfig)
	config.Get(defaultRootPath, "mysql").Scan(&mysqlConfig)

	// 标记已经初始化
	inited = true
}
```

我们目前定义了三个配置结构，它们在basic的[config](user-service/basic/config)目录下

- [profiles](./user-service/basic/config/profiles.go)
- [consul](./user-service/basic/config/config_consul.go)
- [mysql](./user-service/basic/config/config_mysql.go)：

```go
// defaultProfiles 属性配置文件
type defaultProfiles struct {
	Include string `json:"include"`
}

// defaultConsulConfig 默认consul 配置
type defaultConsulConfig struct {
	Enabled bool   `json:"enabled"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
}

// defaultMysqlConfig mysql 配置
type defaultMysqlConfig struct {
	URL               string `json:"url"`
	Enable            bool   `json:"enabled"`
	MaxIdleConnection int    `json:"maxIdleConnection"`
	MaxOpenConnection int    `json:"maxOpenConnection"`
}
```

###### 数据库初始化

数据库的初始化动作在[db.go](user-service/basic/db/db.go)目录下，下面是初始化方法入口：

```go
package db

// ***

var (
	inited  bool
	mysqlDB *sql.DB
	m       sync.RWMutex
)

// InitDB 初始化数据库
func InitDB() {
	m.Lock()
	defer m.Unlock()

	var err error

	if inited {
		err = fmt.Errorf("[initMysql] Mysql 已经初始化过")
		log.Fatal(err)
		return
	}

	// 如果配置声明使用mysql
	if config.GetMysqlConfig().GetEnabled() {
		initMysql()
	}

	inited = true
}
```

从代码中可以看到，在判断配置文件中有激活Mysql指令**GetEnabled**时才会去加载数据库。

[mysql.go](user-service/basic/db/mysql.go)中的初始化代码：

```go
func initMysql() {

	var err error

	// 创建连接
	mysqlDB, err = sql.Open("mysql", config.GetMysqlConfig().GetURL())
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	// 最大连接数
	mysqlDB.SetMaxOpenConns(config.GetMysqlConfig().GetMaxOpenConnection())

	// 最大闲置数
	mysqlDB.SetMaxIdleConns(config.GetMysqlConfig().GetMaxIdleConnection())

	// 激活链接
	if err = mysqlDB.Ping(); err != nil {
		log.Fatal(err)
		panic(err)
	}
}
```

##### 用户模型服务

todo

### user-web

toto

## 延伸阅读

[使用Micro模板新建服务][micro-new]

[micro-new]: https://github.com/micro-in-cn/micro-all-in-one/tree/master/middle-practices/micro-new