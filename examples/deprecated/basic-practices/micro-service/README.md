# Service

Micro中有两种服务类型，一种是就是普通常见的服务，这里就叫service，另一种叫function，它是只响应一次请求的service，响应后便会从注册中心卸载。

## 内容

- [service](./service) 介绍如何使用service
- [function](./function) 介绍如何使用function
- [proto](./proto) 原型文件、类与接口文档

## 预备动作

- 生成proto原型类与接口（已生成，可根据自身情况重新生成）

我们已经预定义好了一个小示例接口原型**greeter.proto**，请切换到其所在目录，然后执行以下命令，该命令会在所在目录下生成类和接口文件。

```bash
$ protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. greeter.proto
```
