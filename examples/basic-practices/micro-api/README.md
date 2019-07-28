# Micro API

本篇介绍如何使用Micro的API gateway。

## 概览

可以先查看[Micro API][什么是API]，有个大概了解，也可以参考[源码][API源码]。我们以下简称**API**。

所以**API**本质上就一个服务网关，它具备动态路由、服务发现的能力，以HTTP方式把外界请求映射到本地微服务，以对外提供服务。

通过服务发现，配合内置的命名空间（Namespace）规则，**API**可以把请求url解析映射到匹配该命名空间规则的服务接口。

在Micro体系中，服务都会有自己的命名空间，而**API**的默认命名空间是**go.micro.api**。通常，部署在其后提供具体接口的微服务，我们会默认按照
Micro风格将它们的命名空间设置为**go.micro.api.example**的样式，**example**便是具体的服务名，如果需要改动这个值，在启动时指定
`--namespace=指定名称`指令即可。

## Handler处理器

Micro API目前有5种处理方式，下面我们会讲到，我们可以根据需求将**API**设置成指定的类型。

| - | 类型 | 说明
----|----|----
1 | rpc | 默认值，通过RPC向go-micro应用转送请求，通常只传送请求body，头信息不封装。只接收POST请求
2 | api | 与rpc差不多，但是会把完整的http头封装向下传送，不限制请求方法
3 | http或proxy | 以反向代理的方式使用**API**，相当于把普通的web应用部署在**API**之后，让外界像调api接口一样调用web服务
4 | web | 与http差不多，但是支持websocket
5 | event | 代理event事件服务类型的请求
6 | meta* | 元数据，通过在代码中的配置选择使用上述中的某一个处理器

- 重点讲一下rpc和api两种类型的区别，它们的区别就在于，rpc不会把请求头信息封装传下去，而api会。
- meta，并无此模式，只是对api或rpc模式下的扩展使用方法。
- 目前版本（V1）无法支持多个handler并存运行

## 请求映射

### RPC/API类型

Micro内部有将http请求路径映射到服务的机制，映射规则可以通过下表介绍

http路径    |    后台服务    |    接口方法
----    |    ----    |    ----
/foo/bar    |    go.micro.api.foo    |    Foo.Bar
/foo/bar/baz    |    go.micro.api.foo    |    Bar.Baz
/foo/bar/baz/cat    |    go.micro.api.foo.bar    |    Baz.Cat

默认的命名空间是**go.micro.api**，上面说过可以通过`--namespace`指令自定义。

而有些带版本号的路径，也可以映射到服务名上

请求路径    |    后台服务    |    接口方法
----    |    ----    |    ----
/foo/bar    |    go.micro.api.foo    |    Foo.Bar
/v1/foo/bar    |    go.micro.api.v1.foo    |    Foo.Bar
/v1/foo/bar/baz    |    go.micro.api.v1.foo    |    Bar.Baz
/v2/foo/bar    |    go.micro.api.v2.foo    |    Foo.Bar
/v2/foo/bar/baz    |    go.micro.api.v2.foo    |    Bar.Baz

从上面的映射规则中可以看出，**RPC/API**模式下，路径后面的两个参数会被组合成Golang公共方法路径名，而剩下的会加上命名空间前缀组成服务名。比如：

`/v1/foo/bar/baz`，其中`bar/baz`首字母大写转成`Bar.Baz`方法路径；剩下的`/v1/foo/`，附加上命名空间前缀`go.micro.api`组成
`go.micro.api.v1.foo`。

### Proxy类型

如果我们启动**API**时传指令`--handler=http`，那么**API**便会反向代理请求到具有API命名空间的后台服务中。

比如：

请求路径    |    服务    |    后台服务路径
---    |    ---    |    ---
/greeter    |    go.micro.api.greeter    |    /greeter
/greeter/:name    |    go.micro.api.greeter    |    /greeter/:name

### Event类型

启动**API**时传指令`--handler=event`，那么**API**便会反向代理请求到具有API命名空间的后台事件消费服务中。

比如（命名空间设置为go.micro.evt）：

请求路径    |    服务    |    方法
---    |    ---    |    ---
/user/login    |    go.micro.evt.user    |    侦听器对象（示例中的new(Event)）所有公共方法，且方法要有ctx和事件参数

## 示例源码

本示例下有6个目录，分别对应**API**的上面说的6种handler处理器类型或模式

[什么是API]: https://micro.mu/docs/cn/api.html
[API源码]: https://github.com/micro/micro/tree/master/api