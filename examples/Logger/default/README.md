# 默认日志

Go-Micro中的默认日志组件并不会打印日志到文件，它只会打印到控制台。

我们运行一下示例代码：

```go
import (
	"github.com/micro/go-micro/v2/logger"
)

func main() {
	logger.Debug("Debug")
	logger.Debugf("Debug %s", "Hello")
	logger.Info("Info")
	logger.Infof("Info %s", "Hello")
	logger.Error("Error")
	logger.Errorf("Debug %s", "Hello")
}
```

打印如下：

```bash
localhost:default micro-cn$ go run main.go 
2020-05-23 22:47:58  level=info Info
2020-05-23 22:47:58  level=info Info Hello
2020-05-23 22:47:58  level=error Error
2020-05-23 22:47:58  level=error Debug Hello
```

大家发现，并没有打印Debug，那里因为默认的级别是Info，我们需要将日志级别调低，有如下两种方式

- 环境变量

设置*MICRO_LOG_LEVEL*

```bash
localhost:default shuxian$ MICRO_LOG_LEVEL=debug go run main.go 
2020-05-23 22:51:52  level=debug Debug
2020-05-23 22:51:52  level=debug Debug Hello
2020-05-23 22:51:52  level=info Info
2020-05-23 22:51:52  level=info Info Hello
2020-05-23 22:51:52  level=error Error
2020-05-23 22:51:52  level=error Debug Hello
```

- 使用接口

```go
logger.Init(logger.WithLevel(logger.DebugLevel))
```

当使用接口时，是可以动态设置日志级别，这方便我们在线上调试：

```go
	logger.Init(logger.WithLevel(logger.DebugLevel))
	logger.Debug("Debug")
	logger.Init(logger.WithLevel(logger.InfoLevel))
	logger.Debug("Debug2")
```

第二个Debug将不会被打印