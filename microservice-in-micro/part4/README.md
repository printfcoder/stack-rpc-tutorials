# 第四章 使用配置中心 doing

前面的章节中，我们把配置以yml文件的方式放到配置目录**conf**中，并通过**go-config**把其加载读到应用中。

Micro生态链中使用**go-config**管理配置，但是目前**go-config**定位是客户端插件，即是说它并不没有充当配置中心的能力，不过，可以通过接口使用其它如consul、etcd、k8s等具备kv存储能力的配置服务器。

**go-config**与配置服务器组合，足以俱备配置中心的能力。

go-config在Micro体系中工作层次如下图所示：

![](../docs/part4_go-config_view.png)

本章我们重点介绍如何使用gRPC和Consul作为配置中心，因为etcd、k8s与consul使用方式大同小异，不过多赘述。

go-config所有使用方式，包括本地与中心服务可参考下列示例

- [Env](https://github.com/micro-in-cn/all-in-one/tree/master/middle-practices/micro-config/env) 本地 基于环境变量
- [File](https://github.com/micro-in-cn/all-in-one/tree/master/basic-practices/micro-config) 本地 基于配置文件
- [Flag](https://github.com/micro-in-cn/all-in-one/tree/master/middle-practices/micro-config/flag) 本地 基于命令行Flag参数文件
- [memory](https://github.com/micro-in-cn/all-in-one/tree/master/middle-practices/micro-config/memory) 本地 基于内存方式配置
- [microcli](https://github.com/micro-in-cn/all-in-one/tree/master/middle-practices/micro-config/microcli) 本地 基于MicroCli参数配置
- [gRPC](https://github.com/micro-in-cn/all-in-one/tree/master/senior-practices/micro-config/gRPC) 使用gRPC服务作为配置中心
- [Consul](https://github.com/micro-in-cn/all-in-one/tree/master/senior-practices/micro-config/consul) 使用Consul服务作为配置中心
- [etcd](https://github.com/micro-in-cn/all-in-one/tree/master/senior-practices/micro-config/consul) 使用etcd服务作为配置中心
- [k8s](https://github.com/micro-in-cn/all-in-one/tree/master/senior-practices/micro-config/consul) 使用K8s服务作为配置中心

## gRPC Server

应用与gRPC Server之间的交互如下

![](../docs/part4_gRPC_server_view.png)



## Consul Server

## 参考阅读

- [如何使用gRPC编写服务](https://medium.com/pantomath/how-we-use-grpc-to-build-a-client-server-system-in-go-dd20045fa1c2)


## 系列文章

- [第一章 用户服务][第一章]
- [第二章 权限服务][第二章]
- [第三章 库存服务、订单服务、支付服务与Session管理][第三章]
- [第五章 消息总线1][第五章] todo
- [第六章 消息总线2][第六章] todo
- [第七章 日志持久化][第四章] todo
- [第八章 熔断、降级、容错与健康检查][第六章] todo
- [第九章 链路追踪][第七章] todo
- [第十章 容器化][第八章] todo
- [第十一章 总结][第十章] todo

[第一章]: ../part1
[第二章]: ../part2
[第三章]: ../part3
[第四章]: ../part4
[第五章]: ../part5
[第六章]: ../part6
[第七章]: ../part7
[第八章]: ../part8
[第九章]: ../part9
[第十章]: ../part10
[第十一章]: ../part11