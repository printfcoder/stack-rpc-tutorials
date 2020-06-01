# Micro 示例集合

## 目录

[老目录](./README.OLD.md)

- [Broker](./broker) micro broker异步消息
 - [basic](./broker/basic)  基础用法
 - [kafka](./broker/kafka)  集成Kafka
 - [nsq](./broker/nsq)  集成Nsq
 - [rabbitMQ](./broker/rabbitmq) 集成RabbitMQ 
- [Client](./client) Go-Micro客户端
  - [json](./client/json) JSON格式调用
  - [retry](./client/retry) 重试
  - [rpc](./client/rpc) RPC调用
  - [selector](./client/selector) 选择器
    - [filter](./client/selector/filter) 过滤器
      - [ip](./client/selector/filter/ip) 指定IP调用
      - [version](./client/selector/filter/version) 指定版本号调用
- [Service](./service) Service的各种用法，Func、Stream（file）、timeout等
  - [Function](./service/function) 函数式服务
  - [Service](./service/service) 常规用法
  - [Stream](./service/stream) 在服务中处理流式数据
  - [Timeout](./service/timeout) 在服务中处理超时
- [Micro API](./micro-api) Micro微服务网关
  - [*rpc-handler](./micro-api/rpc) rpc模式
  - [*api-handler](./micro-api/api) api模式
  - [*proxy-handler](./micro-api/proxy) proxy（http）模式
  - [*web-handler](./micro-api/web) web（websocket）模式
  - [*event-handler](./micro-api/event) event模式
  - [*meta-handler](./micro-api/meta) meta模式
- [Micro Cors](./micro-cors) Micro跨域
- [gRPC](./grpc) gRPC相关用法
