# 第二章 库存服务、订单服务、支付服务与Session管理 doing

我们先回顾下前面两章我们完成的工作

- 第一章 初步完成了用户服务部分的两个子服务**user-web**和**user-service**
- 第二章 完成**auth**服务

待完成工作

- session管理

而本章我们要完成剩下的几个web服务，**orders-web**，**inventory-web**和**payment-web**，以及它们各自对应的服务层应用**orders-service**，**inventory-service**和**payment-service**

下面我们大体介绍下三个web和service服务主要有哪些功能

|服务|接口|---|
|---|---|---|
|orders|---|---|
|inventory|---|---|
|payment|---|---|

[上一章][第二章]我们初步完成了用户服务部分的两个子服务**user-web**和**user-service**。但是最后我们并没有实现session管理，以及抽离公用基础包。

在本篇中，我们除了完成抽离公用基础包，还要实现请求认证服务auth（session管理我们需要放到下一章节来完成，因为现在我们的web服务太少，不方便看效果）。

后面的章节中，user-web，orders-web，inventory-web等接收到的需要认证的请求都要向auth确认。

本章我们要实现**Auth**服务的工作架构如下图：

![](../docs/part2_auth_layer_view.png)

- 当用户请求每个web服务时，会有**wrapper**调用**auth**确定认证结果，并缓存合法结果30分钟。
- 当用户退出时，**auth**广播，各服务**sub**清掉缓存。

我们的缓存使用**redis**。

同时，我们将就在第一章的基础上改动一番，直接把代码复制一份，将import指令中的part1路径换成part2即可。

## 开始写代码

## 总结

## 系列文章

- [第一章 用户服务][第一章]
- [第二章 权限服务][第二章]
- [第四章 消息总线、日志持久化][第四章] todo
- [第五章 使用配置中心][第五章] todo
- [第六章 熔断、降级、容错][第六章] todo
- [第七章 链路追踪][第七章] todo
- [第八章 docker打包与K8s部署][第八章] todo
- [第九章 单元测试][第九章] todo
- [第十章 总结][第十章] todo

## 讨论

朋友，请加入[slack](http://slack.micro.mu/)，进入**中国区**Channel沟通。

## 延伸阅读


[micro-new]: https://github.com/micro-in-cn/all-in-one/tree/master/middle-practices/micro-new
[protoc-gen-go]: https://github.com/micro/protoc-gen-micro
[micro-new-code]: https://github.com/micro/micro/tree/master/new
[go-micro]: https://github.com/micro/go-micro
[go-config]: https://github.com/micro/go-config
[go-web]: https://github.com/micro/go-web
[go-broker]: https://github.com/micro/go-micro/broker
[jwt]: https://jwt.io/introduction/

[第一章]: ../part1
[第二章]: ../part2
[第四章]: ../part4
[第五章]: ../part5
[第六章]: ../part6
[第七章]: ../part7
[第八章]: ../part8
[第九章]: ../part9
[第十章]: ../part10