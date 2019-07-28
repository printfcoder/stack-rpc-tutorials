# Go Config

基础部分我们只讲解传统的文件配置读取，因为这个足够简单，也足够覆盖大多数的使用场景。

Micro Go-config支持的配置文件类型有json、yaml、toml、xml、hcl。

我们挑最常用的json、yaml即可，剩下的三个使用方式大同小异，差异主要在配置文件的内容格式，config内部有策略会自动匹配。

Go-config高阶部分可以参考[go-config进阶][middle-micro-config]

我们不去讨论多个配置文件中冲突的配置会以怎样的机制被覆盖，尽管go-config对同一key多个值的行为有所兼容，但是为什么同样的配置会出现在两个地方，这是应该是部署者需要考虑的。

更多资料请翻阅文档[go-config][go-config]

## 运行程序

```bash
go run main.go
```

[go-config]: https://micro.mu/docs/cn/go-config.html
[middle-micro-config]: ../../middle-practices/micro-config