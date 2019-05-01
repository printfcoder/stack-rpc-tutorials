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

那么，接下来我们聊聊为什么不要。

## 为什么Micro中日志不分级

在Go官方的(log包](https://golang.org/pkg/log/)中，打印日志的方法，并没有根据级别作区分。除非自己实现。更别谈可以像log4j那样的根据实际运行环境只打开某个级别以上日志的指令。

golang官方也提供了相应带级别的日志版本[glog](https://godoc.org/github.com/golang/glog)，它支持如下级别的日志：

- Info
- Warning
- Error
- Fatal

再比如[zap](https://godoc.org/go.uber.org/zap)更丰富

- Info
- Debug
- Error
- Panic
- Warn
- Fatalf
- DPanic
- ...

很明显，为了照顾其它语言的用户，这两个库都或多或少引入了其它语言库的日志风格。

其实日志库应该是要足够简单。

对于debug类型的调试代码不应该保留在生产环境中，它是面向开发者的。

而warning没有谁会关心，因为它不是错误。

至于Info级别，通常都会打印，既然都无论如何都要打印，为什么不直接log呢？

在golang中，通常我们函数的最后一个回参都会是error，而对于error级别的日志，如果我们能够处理这个错误，那说明它不是致命的，既然不是致命的，我们就处理一下，再当成普通日志打出，最后把错误返回调用者。

所以，真正需要的下只有log（logf）方法。

## 如何持久化

我们小篇幅讲解了为什么Micro日志不分级，现在终于要讲如何实现持久化。

在实际的生产环境中，通常日志是要持久化到本地或者其它存储服务中的，比如db、es等

###  

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