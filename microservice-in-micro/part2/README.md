# 第二章 权限服务 doing

[上一章][第一章]我们初步完成了用户服务部分的两个子服务**user-web**和**user-service**。但是最后我们并没有实现session管理，以及有公用基础包没有抽离出来。

在本篇中，我们除了完成刚提的未完成的两个功能外，还要实现请求认证功能auth。为什么我们要把session管理放到auth中，很简单，session会话在其它web服务，如orders-web，inventory-web等中是共享的。

而session会话本质上属于安全的一部分。当用户请求时，我们需要验证请求的合法性，如果用户登录过了，那么我们要判断session是否有该标记，如果有，我们还继续判断用户是否请求非法内容。

那么问题来，我们在第一章中把登录接口放到**user-web**中，当它验证密码完成后，再通知**auth**操作，而**user-web**的其它接口又要先经过**auth**，这样就会造成二者之间耦合，我们要改改，将登录验证操作放到auth中。

所以，我们要在第一章的基础上改动一番，本章我们要实现**Auth**服务的工作架构如下图：

![](../docs/part2_auth_layer_view.png)

- 将**user-web**的接口 **/user/login**移到**auth**服务中。
- 当用户请求每个web服务时，会有**wrapper**调用**auth**确定认证结果，并缓存合法结果30分钟。
- 当用户退出时，**auth**广播，各服务**sub**清掉缓存。

## 开始写代码

使用模板生成**auth**服务代码

```bash
micro new --namespace=mu.micro.book --type=meta --alias=auth github.com/micro-in-cn/tutorials/microservice-in-micro/part2/auth
```



## 总结

## 系列文章

- [第一章 用户服务][第一章]
- [第三章 库存服务与订单服务、支付服务][第三章] todo
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

[第一章]: ../part1
[第三章]: ../part3
[第四章]: ../part4
[第五章]: ../part5
[第六章]: ../part6
[第七章]: ../part7
[第八章]: ../part8
[第九章]: ../part9
[第十章]: ../part10