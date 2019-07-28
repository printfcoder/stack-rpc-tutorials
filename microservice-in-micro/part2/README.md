# 第二章 权限服务

[上一章][第一章]我们初步完成了用户服务部分的两个子服务**user-web**和**user-srv**。但是最后我们并没有实现session管理，以及抽离公用基础包。

在本篇中，我们除了完成抽离公用基础包，还要实现请求认证服务auth（session管理我们需要放到下一章节来完成，因为现在我们的web服务太少，不方便看效果）。

后面的章节中，user-web，orders-web等接收到的需要认证的请求都要向auth确认。

本章我们要实现**Auth**服务的工作架构如下图：

![](../docs/part2_auth_layer_view.png)

- 当用户请求每个web服务时，会有**wrapper**调用**auth**确定认证结果，并缓存合法结果30分钟。
- 当用户退出时，**auth**广播，各服务**sub**清掉缓存。

我们的缓存使用**redis**。

同时，我们将就在第一章的基础上改动一番，直接把代码复制一份，将import指令中的part1路径换成part2即可。

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
├── user-srv
└── user-web
```

然后我们增加**redis**配置与**jwt**配置，其中，jwt属于我们应用自身配置，下面很快我们会讲到，我们会把它放在**app.book**路径下。

[**redis.go**](./basic/config/redis.go)

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

[**config.go**](./basic/config/config.go) 要增加初始化Redis配置的代码

```go
// ...

var (
    // ... 
    redisConfig defaultRedisConfig
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
- 广播token失效

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
service Service {
    rpc MakeAccessToken (Request) returns (Response) {
    }
    rpc DelUserAccessToken (Request) returns (Response) {
    }
}

message Error {
    int32 code = 1;
    string detail = 2;
}

message Request {
    uint64 userId = 1;
    string userName = 2;
    string token = 3;
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

添加配置文件，因为**auth**目前只会用到**consul**、**jwt**、**redis**，故而我们只用添加母文件[application.yml](./auth/conf/application.yml)、[redis](./auth/conf/application-redis.yml)和[consul](./auth/conf/application-consul.yml)配置文件即可。

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
    inited bool
)

// Init 初始化Redis
func Init() {
    m.Lock()
    defer m.Unlock()
    
    if inited {
        log.Log("已经初始化过Redis...")
        return
    }
    
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
    
    // DelUserAccessToken 清除用户token
    DelUserAccessToken(token string) (err error)
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
    
    tokenExpiredTopic = "mu.micro.book.topic.auth.tokenExpired"
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

// DelUserAccessToken 清除用户token
func (s *service) DelUserAccessToken(tk string) (err error) {

    // 解析token字符串
    claims, err := s.parseToken(tk)
    if err != nil {
        return fmt.Errorf("[DelUserAccessToken] 错误的token，err: %s", err)
    }

    // 通过解析到的用户id删除
    err = s.delTokenFromCache(&Subject{
        ID: claims.Subject,
    })

    if err != nil {
        return fmt.Errorf("[DelUserAccessToken] 清除用户token，err: %s", err)
    }

    // 广播删除
    msg := &broker.Message{
        Body: []byte(claims.Subject),
    }
    if err := broker.Publish(tokenExpiredTopic, msg); err != nil {
        log.Logf("[pub] 发布消息失败： %v", err)
    } else {
        fmt.Println("[pub] 发布消息：", string(msg.Body))
    }

    return
}
```

[**access_internal.go**](./auth/model/access/access_internal.go)则是实现内部方法，比如获取缓存key，保存与清除缓存等。

```go
package access

//...

// createTokenClaims Claims
func (s *service) createTokenClaims(subject *Subject) (m *jwt.StandardClaims, err error) {

    now := time.Now()
    m = &jwt.StandardClaims{
        ExpiresAt: now.Add(tokenExpiredDate).Unix(),
        NotBefore: now.Unix(),
        Id:        subject.ID,
        IssuedAt:  now.Unix(),
        Issuer:    "book.micro.mu",
        Subject:   subject.ID,
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

// delTokenFromCache 清空token
func (s *service) delTokenFromCache(subject *Subject) (err error) {
    //保存
    if err = ca.Del(tokenIDKeyPrefix + subject.ID).Err(); err != nil {
        return fmt.Errorf("[delTokenFromCache] 清空token 缓存发生错误，err:" + err.Error())
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

// parseToken 解析token
func (s *service) parseToken(tk string) (c *jwt.StandardClaims, err error) {

    token, err := jwt.Parse(tk, func(token *jwt.Token) (interface{}, error) {
        _, ok := token.Method.(*jwt.SigningMethodHMAC)
        if !ok {
            return nil, fmt.Errorf("不合法的token格式: %v", token.Header["alg"])
        }
        return []byte(config.GetJwtConfig().GetSecretKey()), nil
    })

    // jwt 框架自带了一些检测，如过期，发布者错误等
    if err != nil {
        switch e := err.(type) {
        case *jwt.ValidationError:
            switch e.Errors {
            case jwt.ValidationErrorExpired:
                return nil, fmt.Errorf("[parseToken] 过期的token, err:%s", err)
            default:
                break
            }
            break
        default:
            break
        }

        return nil, fmt.Errorf("[parseToken] 不合法的token, err:%s", err)
    }

    // 检测合法
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        return nil, fmt.Errorf("[parseToken] 不合法的token")
    }

    return mapClaimToJwClaim(claims), nil
}

// 把jwt的claim转成claims
func mapClaimToJwClaim(claims jwt.MapClaims) *jwt.StandardClaims {

    jC := &jwt.StandardClaims{
        Subject: claims["sub"].(string),
    }

    return jC
}
```

<span style="color:red">*</span>对于生产环境的token，光有用户id和用户名是不够的，还需要其他客户端环境信息，比如ip、浏览器指纹、手机MEID等等可以识别客户端在一定范围内唯一性的标识。

这样才能保证token即使泄漏，也极难有可能盗用，我们是为了演示尽可能简单些，省去一些体力活。

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

至此，auth服务基本完成了。下面开始改造**user-web**服务。

### user-web

[**user-web**](./user-web)改动不大，我们只需要改两点

- 把原来的基础包**basic**删除，使用公用包的初始化方法。这一步我们略过，大家直接手动删除即可。
- 改造Login方法，增加获取token逻辑
- 返回带token的set-cookie头信息。
- 增加退出方法Logout

[**handler.go**](./user-web/handler/handler.go)

```go
var (
    serviceClient us.Service
    authClient    auth.Service
)

// Error 错误结构体
type Error struct {
    Code   string `json:"code"`
    Detail string `json:"detail"`
}

func Init() {
    serviceClient = us.NewService("mu.micro.book.srv.user", client.DefaultClient)
    authClient = auth.NewService("mu.micro.book.srv.auth", client.DefaultClient)
}

// Login 登录入口
func Login(w http.ResponseWriter, r *http.Request) {

    // 只接受POST请求
    if r.Method != "POST" {
        log.Logf("非法请求")
        http.Error(w, "非法请求", 400)
        return
    }

    r.ParseForm()

    // 调用后台服务
    rsp, err := serviceClient.QueryUserByName(context.TODO(), &us.Request{
        UserName: r.Form.Get("userName"),
    })
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    // 返回结果
    response := map[string]interface{}{
        "ref": time.Now().UnixNano(),
    }

    if rsp.User.Pwd == r.Form.Get("pwd") {
        response["success"] = rsp.Success

        // 干掉密码返回
        rsp.User.Pwd = ""
        response["data"] = rsp.User
        log.Logf("[Login] 密码校验完成，生成token...")

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

        // 同时将token写到cookies中
        w.Header().Add("set-cookie", "application/json; charset=utf-8")
        // 过期30分钟
        expire := time.Now().Add(30 * time.Minute)
        cookie := http.Cookie{Name: "remember-me-token", Value: rsp2.Token, Path: "/", Expires: expire, MaxAge: 90000}
        http.SetCookie(w, &cookie)

    } else {
        response["success"] = false
        response["error"] = &Error{
            Detail: "密码错误",
        }
    }

    w.Header().Add("Content-Type", "application/json; charset=utf-8")

    // 返回JSON结构
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
}

// Logout 退出登录
func Logout(w http.ResponseWriter, r *http.Request) {

    // 只接受POST请求
    if r.Method != "POST" {
        log.Logf("非法请求")
        http.Error(w, "非法请求", 400)
        return
    }

    tokenCookie, err := r.Cookie("remember-me-token")
    if err != nil {
        log.Logf("token获取失败")
        http.Error(w, "非法请求", 400)
        return
    }

    // 删除token
    _, err = authClient.DelUserAccessToken(context.TODO(), &auth.Request{
        Token: tokenCookie.Value,
    })
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    // 清除cookie
    cookie := http.Cookie{Name: "remember-me-token", Value: "", Path: "/", Expires: time.Now().Add(0 * time.Second), MaxAge: 0}
    http.SetCookie(w, &cookie)

    w.Header().Add("Content-Type", "application/json; charset=utf-8")

    // 返回结果
    response := map[string]interface{}{
        "ref":     time.Now().UnixNano(),
        "success": true,
    }

    // 返回JSON结构
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
}
```

我们在代码中作了如下改动

- 新增了auth客户端**authClient**并在Init中初始化。
- Login 方法中在获取用户信息并比对密码完成后向auth服务申请token，返回时带上token。
- Logout 方法将用户的token发送到auth服务，让auth将token清空

那我们所有的代码都已经写完了，下面我们开始运行程序

运行api

```bash
$ micro --registry=consul --api_namespace=mu.micro.book.web  api --handler=web
```

运行user-srv

```bash
$ cd user-srv
$ go run main.go plugin.go 
```

运行user-web

```bash
$ cd user-web
$ go run main.go
```

运行auth

```bash
$ cd auth
$ go run main.go
```

请求登录

```bash
$ curl --request POST   --url http://127.0.0.1:8080/user/login   --header 'Content-Type: application/x-www-form-urlencoded'  --data 'userName=micro&pwd=1234'
```

返回结果

```json
{
    "data": {
        "id": 10001,
        "name": "micro"
    },
    "ref": 1555428438879936000,
    "success": false,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NTgwMjA0MzgsImp0aSI6IjEwMDAxIiwiaWF0IjoxNTU1NDI4NDM4LCJpc3MiOiJib29rLm1pY3JvLm11IiwibmJmIjoxNTU1NDI4NDM4LCJzdWIiOiIxMDAwMSJ9._OCu2umIFuFmNwn-sSHoXBrjyBovhgcPQweOTOlZqq8"
}
```

退出登录（需要将token换成实际的）

```bash
$ curl --request POST \
  --url http://127.0.0.1:8080/user/logout \
  --cookie 'remember-me-token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NTgwMjA0MzgsImp0aSI6IjEwMDAxIiwiaWF0IjoxNTU1NDI4NDM4LCJpc3MiOiJib29rLm1pY3JvLm11IiwibmJmIjoxNTU1NDI4NDM4LCJzdWIiOiIxMDAwMSJ9._OCu2umIFuFmNwn-sSHoXBrjyBovhgcPQweOTOlZqq8'
```

## 总结

本章我们在第一章的基础上增加认证服务auth，并优化了**user-web**和**user-srv**服务，把基础组件抽到公用包中。

我们本章用到的Micro技术点有（依次从文章开始到结束）

- [**micro new**][micro-new-code]，生成Micro风格的模板代码，它是micro项目中的一个子包。
- [**protoc-gen-go**][protoc-gen-go]，隐藏在`protoc ... --micro_out`指令中执行了，感兴趣的同学可以去了解一下。
- [**go-micro**][go-micro]，代码中**micro.NewService**，**service.Init**等都是go-micro中不同类型服务各自实现的方法。
- [**go-config**][go-config]，加载配置时使用。
- [**go-web**][go-web]，编写web应用**user-web**时用到。
- [**go-broker**][go-broker]，编写web应用**user-web**时用到。

至此，我们的服务架构完成度如下：（红圈服务为完成）

![](../docs/part2_done_services.png)

但我们工作还不完善：

- 仍然遗留session管理问题。
- 重复登录、退出不会报错，不过这也不是这里要考虑的，太过度关注这些细节会让这系列教程演变得非常重，没有重点。

接下来的下一章，我们会编写剩下的web服务，刚提到的session管理，我们会在这一章实现。

请翻阅，[第三章 库存服务、订单服务、支付服务与Session管理][第三章]。

## 系列文章

- [第一章 用户服务][第一章]
- [第三章 库存服务、订单服务、支付服务与Session管理][第三章]
- [第四章 使用配置中心][第四章]
- [第五章 日志持久化][第五章]
- [第六章 熔断、降级、容错与健康检查][第六章]
- [第七章 链路追踪][第七章]
- [第八章 容器化][第八章]

## 讨论

朋友，请加入[slack](http://slack.micro.mu/)，进入**中国区**Channel沟通。

## 延伸阅读

[micro-new]: https://github.com/micro-in-cn/tutorials/examples/tree/master/middle-practices/micro-new
[protoc-gen-go]: https://github.com/micro/protoc-gen-micro
[micro-new-code]: https://github.com/micro/micro/tree/master/new
[go-micro]: https://github.com/micro/go-micro
[go-config]: https://github.com/go-micro/config
[go-web]: https://github.com/micro/go-web
[go-broker]: https://github.com/micro/go-micro/broker
[jwt]: https://jwt.io/introduction/

[第一章]: ../part1
[第三章]: ../part3
[第四章]: ../part4
[第五章]: ../part5
[第六章]: ../part6
[第七章]: ../part7
[第八章]: ../part8
