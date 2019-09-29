# micro 最佳实践

本项目旨在以全面的视角向大家演示micro服务的开发，并尽可能介绍其功能。我们尝试由浅入深，从基础篇开始，每个目录对应一个级别的示例。

## 讨论

朋友，请加入[slack](http://slack.micro.mu/)，进入**中国区**Channel沟通。

## Micro源码运行环境安装

这里省略了Golang环境安装，未安装的同学可参考[Golang][golang-cn]，选择合适自己的安装包搭建Go开发环境。

### 安装micro

```bash
$ go get -u github.com/micro/micro
```

## 目录

*代表是初阶，**代表中阶，***代表高阶

- [micro-api](./basic-practices/micro-api) 
  - [*rpc-handler](./basic-practices/micro-api/rpc) rpc模式
  - [*api-handler](./basic-practices/micro-api/api) api模式
  - [*proxy-handler](./basic-practices/micro-api/proxy) proxy（http）模式
  - [*web-handler](./basic-practices/micro-api/web) web（websocket）模式
  - [*event-handler](./basic-practices/micro-api/event) event模式
  - [*meta-handler](./basic-practices/micro-api/meta) meta模式
- [micro-broker](./basic-practices/micro-broker) 
  - [*发布与订阅](./basic-practices/micro-broker/basic) 基础篇
  - [*Kafka](./basic-practices/micro-broker/kafka) 使用Kafka作为消息中间件
  - [*NATs](./basic-practices/micro-broker/nats) 使用NATs作为消息中间件
  - [*nsq](./basic-practices/micro-broker/nsq) 使用NSQ作为消息中间件
  - [*rabbitMQ](./basic-practices/micro-broker/rabbitmq) 使用rabbitMQ作为消息中间件
- [micro-cli](./middle-practices/micro-cli) 如何使用命令行接口
  - [**flag](./middle-practices/micro-cli/flags) 如何使用flag
- [micro-client](./middle-practices/micro-client) 如何使用Micro客户端
  - [**JSON](./middle-practices/micro-client/json) JSON客户端
  - [**Retry](./middle-practices/micro-client/retry) 客户端重试
- micro-config 如何读取配置
  - local  本地
    - [*file](./basic-practices/micro-config) 基于文件配置
    - [**env](./basic-practices/micro-config) 使用环境变量配置
    - [**flag](./basic-practices/micro-config) 使用命令行参数配置
    - [**memory](./basic-practices/micro-config) 使用环境变量配置
    - [**microcli](./basic-practices/micro-config) 使用MicroCli指令参数配置
  - [***center](./senior-practices/micro-config) 使用配置中心 todo
    - [***consul](./senior-practices/micro-config/consul) 使用consul作为配置中心 todo
    - [***etcd](./senior-practices/micro-config/etcd) 使用etcd作为配置中心 todo
    - [***gRPC](./senior-practices/micro-config/gRPC) 使用gRPC作为配置中心
    - [***k8s](./senior-practices/micro-config/k8s) 使用k8s作为配置中心 todo
- [micro-grpc](./middle-practices/micro-grpc) micro与grpc相关用法
- [**micro-new](./middle-practices/micro-new) 新建模板
- [*micro-registry](./basic-practices/micro-registry) 
  - [*etcdv3](./basic-practices/micro-registry/etcdv3)
- [micro-service](./basic-practices/micro-service) 编写Micro服务，包含service和function
  - [*function](./basic-practices/micro-service/function)
  - [*service](./basic-practices/micro-service/service)
  - [**自定义broker](./middle-practices/micro-service/custom-broker)
  - [**关闭应用](./middle-practices/micro-service/shutdown)
- [micro-sync](./senior-practices/micro-sync) 分布式同步、锁、时间、任务等
  - [***consul-lock](./senior-practices/micro-sync/lock/consul) 基于consul的锁
- [micro-client-filter](./senior-practices/micro-filter) 客户端过滤器
  - [***version-filter](./senior-practices/micro-filter/version) 版本号过滤器
- [micro-web](./basic-practices/micro-web) 
  - [*web-pub](./basic-practices/micro-web/web-pub) web与pub结合
- [**micro-wrapper](./middle-practices/micro-wrapper) 
  - [**基础用法](./middle-practices/micro-wrapper/basic)
- grpc
  - [**grpc-oneof**](./middle-practices/grpc-oneof) grpc oneOf 示例todo

[golang-cn]: https://golang.google.cn/
[tutorials]: https://github.com/micro-in-cn/tutorials
