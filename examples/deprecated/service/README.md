# Service

Micro中有两种服务类型，一种是就是普通常见的服务，这里就叫service，另一种叫function，它是只响应一次请求的service，响应后便会从注册中心卸载。

## 内容

- [Function](./function) 函数式服务
- [Service](./service) 常规用法
- [Stream](./stream) 在服务中处理流式数据
- [Timeout](./timeout) 在服务中处理超时

## 预备动作

- 生成proto原型类与接口（已生成，可根据自身情况重新生成）

我们已经预定义好了一个小示例接口原型**greeter.proto**，请切换到其所在目录，然后执行以下命令，该命令会在所在目录下生成类和接口文件。

```bash
$ protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. proto/greeter.proto
```
