# 第七章 链路追踪

微服务系统中，客户端的一次请求，可能需要经过系统中多个模块的相互协作才能完成。为了追踪客户端的一次操作背后调用了哪些模块以及调用先后顺序，我们需要一个分布式链路追踪系统。我们选择 Jaeger，Jaeger 是 Uber 推出的一款开源分布式追踪系统，兼容 OpenTracing API。

OpenTracing 通过提供平台无关、厂商无关的API，使得开发人员能够方便的添加（或更换）追踪系统的实现。 OpenTracing 提供了用于运营支撑系统的和针对特定平台的辅助程序库，所以我们使用 OpenTracing 的程序库可以方便的替换追踪工具。 

## 部署Jaeger实验环境

Jaeger 提供一个all in one 的docker镜像，可以快速搭建实验环境

```bash
docker run -d --name jaeger 
-e COLLECTOR_ZIPKIN_HTTP_PORT=9411 
-p 5775:5775/udp 
-p 6831:6831/udp 
-p 6832:6832/udp 
-p 5778:5778 
-p 16686:16686 
-p 14268:14268 
-p 9411:9411 
jaegertracing/all-in-one:1.6
```

然后，打开  http://localhost:16686 进入 Jaeger 的 UI

## micro链路追踪插件
### micro自带的opentracing插件

在micro自带的插件中已经有opentracing的插件了，包含server，client等，在 **micro/go-plugins/wrapper/trace/opentracing** 目录下，不过这个插件只能go-micro构建的微服务(api，srv)中使用。

我们可以在构建服务的时候直接使用，只需要在服务初始化时增加一行函数就可以了。

```go
service := micro.NewService(
        micro.Name(name),
        micro.Version("latest"),
        micro.WrapHandler(opentracing.NewHandlerWrapper(opentracing.GlobalTracer())),
    )
```

### 为micro API 网关增加插件



## 总结



## 参考阅读

[分布式集群环境下调用链路追踪](https://www.ibm.com/developerworks/cn/web/wa-distributed-systems-request-tracing/index.html)

[OpenTracing中文文档](https://wu-sheng.gitbooks.io/opentracing-io/content/)
## 系列文章

- [第一章 用户服务][第一章]
- [第二章 权限服务][第二章]
- [第三章 库存服务、订单服务、支付服务与Session管理][第三章]
- [第四章 使用配置中心][第四章]
- [第五章 日志持久化][第五章]
- [第六章 熔断、降级、容错与健康检查][第六章]
- [第八章 容器化][第八章] todo
- [第九章 总结][第九章] todo

[第一章]: ../part1
[第二章]: ../part2
[第三章]: ../part3
[第四章]: ../part4
[第五章]: ../part5
[第六章]: ../part6
[第七章]: ../part7
[第八章]: ../part8
[第九章]: ../part9