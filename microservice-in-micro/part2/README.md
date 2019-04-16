# 第二章 权限服务 doing

[上一章][第一章]我们初步完成了用户服务部分的两个子服务**user-web**和**user-service**。但是最后我们并没有实现session管理，以及有公用基础包抽离。

在本篇中，我们除了完成刚提的未完成的两个功能外，还要实现请求认证服务auth。

user-web，orders-web，inventory-web等需要认证的请求都要向auth确认。

所以，我们要在第一章的基础上改动一番，本章我们要实现**Auth**服务的工作架构如下图：

![](../docs/part2_auth_layer_view.png)

- 优化**user-web**的接口，返回带token的set-cookies。
- 当用户请求每个web服务时，会有**wrapper**调用**auth**确定认证结果，并缓存合法结果30分钟。
- 当用户退出时，**auth**广播，各服务**sub**清掉缓存。

我们的缓存使用**redis**。

## 开始写代码

### 优化公用包

在第一章我们提到过要优化配置与基础组件部分的代码，将它们抽取到公用部分。

那我们现在开始。

新建一个**basic**目录用于存放公用代码，先将基础组件代码移到其中。

```text
├── README.md
├── basic
│   ├── basic.go
│   ├── config
│   │   ├── config.go
│   │   ├── config_consul.go
│   │   ├── config_mysql.go
│   │   └── profiles.go
│   └── db
│       ├── db.go
│       └── mysql.go
├── docs
├── user-service
└── user-web
```

然后我们增加**redis**配置与**jwt**配置，其中，jwt属于我们应用自身配置，下面很快我们会讲到，我们会把它放在**app.book**路径下。

[**redis.go**](./basic/config/redis.go)我们也要改动部分

```go
// ...

// RedisConfig redis 配置
type RedisConfig interface {
	GetEnabled() bool
	GetConn() string
	GetPassword() string
	GetDBNum() int
	GetSentinelConfig() RedisSentinelConfig
}

// RedisSentinelConfig 哨兵配置
type RedisSentinelConfig interface {
	GetEnabled() bool
	GetMaster() string
	GetNodes() []string
}

// ...
```

[**config.go**](./basic/config/config.go)我们也要改动部分

```go
// ...

var (
	// ...
	redisConfig             defaultRedisConfig
)

// InitConfig 初始化配置
func InitConfig() {
	
    // redis
	config.Get(defaultRootPath, "redis").Scan(&redisConfig)

}
// ...

// GetRedisConfig 获取Consul配置
func GetRedisConfig() (ret RedisConfig) {
	return redisConfig
}
```

redis配置代码过长，这里我们只贴公有方法部分。

[**jwt**](./basic/config/jwt.go)代码类似，这里不赘述。

下面开始编写**auth**服务。

### auth

auth服务目前需要具备以下能力

- 加载配置
- 生成与验证token

那我们还需要给auth增加token生成策略，因此，我们引入[jwt][jwt]。

jwt是JSON Web Token的简称，它是web服务token安全验证的高效解决方案。本项目中，我们用它生成token与验证token。

jwt Token属于自包含的token，它结构有三个部分组成

- Header 键值对，声明token签名算法，类型等元数据信息
- Payload 键值对，里面一般存放用户数据
- Signature 由**Header+Payload+密钥**加密生成

三者会被Base64-URL各自编码成字符串，再组合成`xxxxx.yyyyy.zzzzz`的形式。

验证的过程就是重新再把**Header+Payload+密钥**加密，看二者的Signature是否一致。

下引我们安装jwt的golang库和redis库，jwt生成后我们会放到redis中

```bash
go get -u github.com/dgrijalva/jwt-go
go get -u github.com/go-redis/redis

```

使用模板生成**auth**服务代码

```bash
micro new --namespace=mu.micro.book --type=srv --alias=auth github.com/micro-in-cn/tutorials/microservice-in-micro/part2/auth
```

把**proto/example/example.proto**文件改成下面的样子，并将文件名与目录改成*auth*，让其能正确生成我们需要的类型与接口

```proto
syntax = "proto3";

package mu.micro.book.srv.auth;

service Service {
    rpc MakeAccessToken (Request) returns (Response) {
    }
}

message Error {
    int32 code = 1;
    string detail = 2;
}

message Request {
    uint64 userId = 1;
    string userName = 2;
}

message Response {
    bool success = 1;
    Error error = 2;
    string token = 3;
}
```

生成原型文件：

```bash
cd auth
protoc --proto_path=. --go_out=. --micro_out=. proto/auth/auth.proto
```

添加配置文件，因为**auth**目前只会用到**consul**、**jwt**、**redis**，故而我们只用添加母文件[application.yml](./auth/conf/application.yml)和[redis](./auth/conf/application-redis.yml))配置文件即可。

application.yml，我们把jwt配置加到其中。

```yaml
app:
  jwt:
    secretKey: W6VjDud2W1kMG3BicbMNlGgI4ZfcoHtMGLWr
```

application-redis.yml，我们加上了哨兵模式的配置，但是我们不会用到哨兵，这里只是给大家演示可以这么声明配置

```yaml
app:
  redis:
    enabled: true
    conn: 127.0.0.1:6379
    dbNum: 8
    password:
    timeout: 3000
    sentinel:
      enabled: false
      master: bookMaster
      nodes: 127.0.0.1:16379,127.0.0.1:26379,127.0.0.1:36379
```

下面我们编写初始化Redis数据库的**Init**方法

[**redis.go**](./basic/redis/redis.go)

```go
package redis

/// ...
var (
	client *redis.Client
	m      sync.RWMutex
)

// Init 初始化Redis
func Init() {
	m.Lock()
	defer m.Unlock()

	redisConfig := config.GetRedisConfig()

	// 打开才加载
	if redisConfig != nil && redisConfig.GetEnabled() {

		log.Log("初始化Redis...")

		// 加载哨兵模式
		if redisConfig.GetSentinelConfig() != nil && redisConfig.GetSentinelConfig().GetEnabled() {
			log.Log("初始化Redis，哨兵模式...")
			initSentinel(redisConfig)
		} else { // 普通模式
			log.Log("初始化Redis，普通模式...")
			initSingle(redisConfig)
		}

		log.Log("初始化Redis，检测连接...")

		pong, err := client.Ping().Result()
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Log("初始化Redis，检测连接Ping.")
		log.Log("初始化Redis，检测连接Ping..")
		log.Logf("初始化Redis，检测连接Ping... %s", pong)
	}
}

// GetRedis 获取redis
func GetRedis() *redis.Client {
	return client
}

func initSentinel(redisConfig config.RedisConfig) {
	client = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    redisConfig.GetSentinelConfig().GetMaster(),
		SentinelAddrs: redisConfig.GetSentinelConfig().GetNodes(),
		DB:            redisConfig.GetDBNum(),
		Password:      redisConfig.GetPassword(),
	})

}

func initSingle(redisConfig config.RedisConfig) {
	client = redis.NewClient(&redis.Options{
		Addr:     redisConfig.GetConn(),
		Password: redisConfig.GetPassword(), // no password set
		DB:       redisConfig.GetDBNum(),    // use default DB
	})
}
```

我们新建目录[model](./auth/model)，将生成逻辑放到model包下，因为它不属于handler接口层需要处理的逻辑，将其封装到服务层或叫业务模型层。

它的结构如下：

```text
...
├── model
│   ├── access             # 用户操作行为相关包
│   │   ├── access.go      # 负责定义、初始化等
│   │   ├── access_internal.go # 内部类，包含保存、清除缓存逻辑
│   │   └── access_token.go    # 生成与获取token的主要代码
│   └── model.go           # 业务模型初始化入口
...
```

开始编写生成token与将其保存到redis的代码。

[**access.go**](./auth/model/access/access.go)

```go
package access

// ...

var (
	s  *service
	ca *r.Client
	m  sync.RWMutex
)

// service 服务
type service struct {
}

// Service 用户服务类
type Service interface {
	// MakeAccessToken 生成token
	MakeAccessToken(subject *Subject) (ret string, err error)

	// GetCachedAccessToken 获取缓存的token
	GetCachedAccessToken(subject *Subject) (ret string, err error)
}

// GetService 获取服务类
func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetService] GetService 未初始化")
	}
	return s, nil
}

// Init 初始化用户服务层
func Init() {
	m.Lock()
	defer m.Unlock()

	if s != nil {
		return
	}

	ca = redis.GetRedis()

	s = &service{}
}
```

我们定义了**service**服务结构，它提供两个方法：

- MakeAccessToken 生成Token
- GetCachedAccessToken 获取缓存的Token

与其它服务类一样，入口方法都是**Init**，获取服务的方法都是**GetService**。

接下来是生成与获取token的方法，它负责实现接口**Service**的两个方法**MakeAccessToken**和**GetCachedAccessToken**。

[**access_token.go**](./auth/model/access/access_token.go)

```go
package access

// ...

var (
	// tokenExpiredDate app token过期日期 30天
	tokenExpiredDate = 3600 * 24 * 30 * time.Second

	// tokenIDKeyPrefix tokenID 前缀
	tokenIDKeyPrefix = "token:auth:id:"
)
// Subject token 持有者
type Subject struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
}

// standardClaims token 标准的Claims
type standardClaims struct {
	SubjectID string `json:"subjectId,omitempty"`
	Name      string `json:"name,omitempty"`
	jwt.StandardClaims
}

// MakeAccessToken 生成token并保存到redis
func (s *service) MakeAccessToken(subject *Subject) (ret string, err error) {

	m, err := s.createTokenClaims(subject)
	if err != nil {
		return "", fmt.Errorf("[MakeAccessToken] 创建token Claim 失败，err: %s", err)
	}

	// 创建
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, m)
	ret, err = token.SignedString([]byte(config.GetJwtConfig().GetSecretKey()))
	if err != nil {
		return "", fmt.Errorf("[MakeAccessToken] 创建token失败，err: %s", err)
	}

	// 保存到redis
	err = s.saveTokenToCache(subject, ret)
	if err != nil {
		return "", fmt.Errorf("[MakeAccessToken] 保存token到缓存失败，err: %s", err)
	}

	return
}

// GetCachedAccessToken 获取token
func (s *service) GetCachedAccessToken(subject *Subject) (ret string, err error) {
	ret, err = s.getTokenFromCache(subject)
	if err != nil {
		return "", fmt.Errorf("[GetCachedAccessToken] 从缓存获取token失败，err: %s", err)
	}

	return
}
```

[**access_internal.go**](./auth/model/access/access_internal.go)则是实现内部方法，比如获取缓存key，保存与清除缓存等。

```go
package access

//...

// createTokenClaims Claims
func (s *service) createTokenClaims(subject *Subject) (m *standardClaims, err error) {

	now := time.Now()
	m = &standardClaims{
		SubjectID: subject.ID,
		Name:      subject.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(tokenExpiredDate).Unix(),
			NotBefore: now.Unix(),
			Id:        subject.ID,
			IssuedAt:  now.Unix(),
			Issuer:    "book.micro.mu",
			Subject:   subject.ID,
		},
	}

	return
}

// saveTokenToCache 保存token到缓存
func (s *service) saveTokenToCache(subject *Subject, val string) (err error) {
	//保存
	if err = ca.Set(tokenIDKeyPrefix+subject.ID, val, tokenExpiredDate).Err(); err != nil {
		return fmt.Errorf("[saveTokenToCache] 保存token到缓存发生错误，err:" + err.Error())
	}
	return
}

// clearTokenFromCache 清空token
func (s *service) clearTokenFromCache(subject *Subject) (err error) {
	//保存
	if err = ca.Del(tokenIDKeyPrefix + subject.ID).Err(); err != nil {
		return fmt.Errorf("[clearTokenFromCache] 清空token 缓存发生错误，err:" + err.Error())
	}
	return
}

// getTokenFromCache 从缓存获取token
func (s *service) getTokenFromCache(subject *Subject) (token string, err error) {

	// 获取
	tokenCached, err := ca.Get(tokenIDKeyPrefix + subject.ID).Result()
	if err != nil {
		return token, fmt.Errorf("[getTokenFromCache] token不存在 %s", err)
	}

	return string(tokenCached), nil
}

```

下面我们改造**auth**的main方法，让其能加载配置：

```go
package main

// ...

func main() {

	// 初始化配置、数据库等信息
	basic.Init()

	// 使用consul注册
	micReg := consul.NewRegistry(registryOptions)

	// 新建服务
	service := micro.NewService(
		micro.Name("mu.micro.book.srv.auth"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)

	// 服务初始化
	service.Init(
		micro.Action(func(c *cli.Context) {
			// 初始化handler
            model.Init()
            // 初始化handler
            handler.Init()
		}),
	)

	// 注册服务
	s.RegisterAuthHandler(service.Server(), new(handler.Service))

	// 启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
// ...
```

同样，在main入口立即初始化基础组件，在Action中初始化业务组件。

至此，auth服务基本完成了。下面开始改造**user-web**方法。

### user-web

[**user-web**](./user-web)改动不大，我们只需要改两点

- 把原来的基础包**basic**删除，使用公用包的初始化方法。这一步我们略过，大家直接手动删除即可。
- 改造Login方法，增加获取token逻辑

[**handler.go**](./user-web/handler/handler.go)

```go
package handler

import (
    // ...

	auth "github.com/micro-in-cn/tutorials/microservice-in-micro/part2/auth/proto/auth"
	//...
)

var (
	// ...
	authClient    auth.Service
)

func Init() {
	// ...
	authClient = auth.NewService("mu.micro.book.srv.auth", client.DefaultClient)
}

// Login 登录入口
func Login(w http.ResponseWriter, r *http.Request) {

	// ...
	// 返回结果
	response := map[string]interface{}{
		"ref": time.Now().UnixNano(),
	}

	// ...
		// 生成token
		rsp2, err := authClient.MakeAccessToken(context.TODO(), &auth.Request{
			UserId:   rsp.User.Id,
			UserName: rsp.User.Name,
		})
		if err != nil {
			log.Logf("[Login] 创建token失败，err：%s", err)
			http.Error(w, err.Error(), 500)
			return
		}

		log.Logf("[Login] token %s", rsp2.Token)
		response["token"] = rsp2.Token

     // ...
}
```

我们新增了auth客户端**authClient**并在Init中初始化。Login方法中在获取用户信息并比对密码完成后向auth服务申请token，返回时带上token。

### session管理



## 总结

## 系列文章

- [第一章 用户服务][第一章]
- [第三章 库存服务、订单服务、支付服务与Session管理][第三章] todo
- [第四章 消息总线、日志持久化][第四章] todo
- [第五章 使用配置中心][第五章] todo
- [第六章 熔断、降级、容错][第六章] todo
- [第七章 链路追踪][第七章] todo
- [第八章 docker打包与K8s部署][第八章] todo
- [第九章 单元测试][第九章] todo
- [第十章 总结][第十章] todo

## 延伸阅读


[micro-new]: https://github.com/micro-in-cn/all-in-one/tree/master/middle-practices/micro-new
[protoc-gen-go]: https://github.com/micro/protoc-gen-micro
[micro-new-code]: https://github.com/micro/micro/tree/master/new
[go-micro]: https://github.com/micro/go-micro
[go-config]: https://github.com/micro/go-config
[go-web]: https://github.com/micro/go-web
[jwt]: https://jwt.io/introduction/

[第一章]: ../part1
[第三章]: ../part3
[第四章]: ../part4
[第五章]: ../part5
[第六章]: ../part6
[第七章]: ../part7
[第八章]: ../part8
[第九章]: ../part9
[第十章]: ../part10