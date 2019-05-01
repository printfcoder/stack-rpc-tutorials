# 第五章 日志持久化

Micro的日志插件[go-log](https://github.com/micro/go-log)并没有提供日志持久化方案，它只一个基于github.com/go-log/log的日志打印接口。

我们先看看**go-log**提供的接口

```go
var (
	// the local logger
	logger log.Logger = golog.New()
)

// Log makes use of github.com/go-log/log.Log
func Log(v ...interface{}) {
	logger.Log(v...)
}

// Logf makes use of github.com/go-log/log.Logf
func Logf(format string, v ...interface{}) {
	logger.Logf(format, v...)
}

// Fatal logs with Log and then exits with os.Exit(1)
func Fatal(v ...interface{}) {
	Log(v...)
	os.Exit(1)
}

// Fatalf logs with Logf and then exits with os.Exit(1)
func Fatalf(format string, v ...interface{}) {
	Logf(format, v...)
	os.Exit(1)
}

// SetLogger sets the local logger
func SetLogger(l log.Logger) {
	logger = l
}
```

源码非常简单，关于打印的总得来说两个接口**Log**和**Fatal**。一个负责打印普通日志，一个则打印致命日志。

还有一个**SetLogger**，则是用来设置默认的日志处理器的。有朋友可能就会产生疑问，特别是从Java转过来的伙伴，为什么没有debug/info/warn/error等常见的日志级别方法。

接下来我们聊聊为什么不要，


在实际的生产环境中，通常日志是要持久化到本地或者其它服务中的。

## 参考阅读

[聊聊日志打印](https://dave.cheney.net/2015/11/05/lets-talk-about-logging)

## 系列文章

- [第一章 用户服务][第一章]
- [第二章 权限服务][第二章]
- [第三章 库存服务、订单服务、支付服务与Session管理][第三章]
- [第四章 使用配置中心][第四章]
- [第六章 熔断、降级、容错与健康检查][第六章] todo
- [第七章 链路追踪][第七章] todo
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