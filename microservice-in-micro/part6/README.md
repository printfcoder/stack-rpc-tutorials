# 第六章 熔断、降级、容错与健康检查

## 熔断、降级、容错
容错，字面意思就是可以容下错误，不让错误再次扩张，让这个错误产生的影响在一定范围之内。我们常见的降级，限流，熔断器，超时重试等等都是容错的方法。服务降级通常指业务上的降级，比如淘宝双十一为了承受巨大的流量会暂时关闭退款功能。当然熔断也是属于服务降级的一种策略，当一个服务出现了故障，直接熔断这个服务，而不是等待调用超时。

我们将使用Hystrix来演示熔断在micro的使用方法。Hystrix的golang的[实现库](https://github.com/afex/hystrix-go):
```Bash
go get github.com/afex/hystrix-go
```
### 如何使用Hystrix
使用Hystrix我们需要根据具体情况配置一些参数，才能达到使用情况。
```go
// DefaultTimeout is how long to wait for command to complete, in milliseconds
DefaultTimeout = 1000
// DefaultMaxConcurrent is how many commands of the same type can run at the same time
DefaultMaxConcurrent = 10
// DefaultVolumeThreshold is the minimum number of requests needed before a circuit can be tripped due to health
DefaultVolumeThreshold = 20
// DefaultSleepWindow is how long, in milliseconds, to wait after a circuit opens before testing for recovery
DefaultSleepWindow = 5000
// DefaultErrorPercentThreshold causes circuits to open once the rolling measure of errors exceeds this percent of requests
DefaultErrorPercentThreshold = 50
```
VolumeThreshold就是单位时间内(10s)触发熔断的最低请求次数，默认为20。ErrorPercentThreshold即触发熔断要达到的错误率，默认为50%。总结一下，即在单位时间内如果调用次数超过20次，且错误率超过50%就触发熔断。

熔断打开之后，再有请求调用的时候，将不会调用主逻辑，而是直接调用降级逻辑。这个时候就会快速返回，而不是等超时后才返回。

在熔断器打开之后，Hystrix会启动一个休眠时间窗SleepWindow，默认为5s。当休眠时间窗到期，熔断器就进入半开状态，允许一次请求到原来的主逻辑上。如果此次请求正常返回，那么熔断器将会关闭，主逻辑恢复正常。否则，熔断器继续保持打开状态，而休眠时间窗会重新计时。

关于Hystrix的详细实现原理，可以查看官方[wiki](https://github.com/Netflix/Hystrix/wiki)。

Hystrix其实就是包装客户端的RPC调用，在这个包装层里实现熔断策略。hystrix-go提供同步和异步两种方式来包装RPC调用。在go-plugins/wrapper/breaker/hystrix中micro已有一个现成插件：
```go
package hystrix

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/client"

	"context"
)

type clientWrapper struct {
	client.Client
}

func (c *clientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	return hystrix.Do(req.Service()+"."+req.Endpoint(), func() error {
		return c.Client.Call(ctx, req, rsp, opts...)
	}, nil)
}

// NewClientWrapper returns a hystrix client Wrapper.
func NewClientWrapper() client.Wrapper {
	return func(c client.Client) client.Client {
		return &clientWrapper{c}
	}
}
```
其中，hystrix.Do方法传入的第三个参数是nil,我们来看看方法的声明
```go
type fallbackFunc func(error) error
...
func Do(name string, run runFunc, fallback fallbackFunc)
```
第三个参数fallback传入一个函数，当主逻辑run方法执行失败或者熔断器触发了，就会执行fallback函数。

修改每个web服务handler包的Init方法，为客户端增加Hystrix插件
```go
import (
...
	"github.com/micro/go-plugins/wrapper/breaker/hystrix"
...
)

func Init() {
	client.DefaultClient.Init(
		client.Wrap(hystrix.NewClientWrapper()),
	)
...
}
```


## 健康检查
在微服务架构中，每个服务都会存在多个实例，可能部署在不同的主机中。因为有网络或者主机等不确定因素，所以每个服务都可能会出现故障。我们需要能够掌握每个服务实例的健康状态，当一个服务故障时，及时将它从注册中心删除。

micro服务注册功能提供的两个参数来实现健康检查功能：
```
micro.RegisterTTL(time.Second*30)
micro.RegisterInterval(time.Second*20)
```
Interval就是间隔多久服务会自动重新注册，TTL就是注册服务的过期时间，如果服务超过过期时间没有去重新注册，注册中心会将其删除。

其实现在go-micro/server/rpc_server.go中,micro使用一个定时器按照设定的间隔时间去自动重新注册。当服务意外故障，无法向注册中心重新注册时，如果超过了设定的TTL时间，注册中心就会将服务删除。

![](../docs/part6_register_code.png)

修改每个服务的main函数，增加两行代码，srv服务如下：
```
...
// 新建服务
	service := micro.NewService(
		micro.Name("mu.micro.book.srv.user"),
+		micro.RegisterTTL(time.Second*15),
+		micro.RegisterInterval(time.Second*10),
		micro.Registry(micReg),
		micro.Version("latest"),
	)
...
```
web服务如下：
```
...
	// 创建新服务
	service := web.NewService(
		web.Name(cfg.Name),
		web.Version(cfg.Version),
+		web.RegisterTTL(time.Second*15),
+		web.RegisterInterval(time.Second*10),
		web.Registry(micReg),
		web.Address(cfg.Addr()),
	)
...

```