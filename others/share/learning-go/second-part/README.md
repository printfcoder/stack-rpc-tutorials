# 夜读第二期

- 视频里面是针对v1的,本文档是针对v2的.
- 以下操作默认路径是github.com/micro-in-cn/tutorials/others/share/learning-go/second-part.

## 启动各个服务

- 启动Sum-srv：

```shell
go run sum-srv/main.go
```

如果运行 go run sum-srv/main.go -h 会看到

```info
--learning_go value                  help一下，你就知道
```

- 启动Prime-src

```shell
go run prime-srv/main.go  
```


- 启动Portal-web

```shell
go run portal-web/main.go
```

> 在浏览器访问:http://localhost:8088/learning/sum?input=10 得到55的结果,且Sum-srv的有一条日志

- 启动log-src

```shell
go run log-srv/main.go
```

> 在浏览器访问:http://localhost:8088/learning/sum?input=10, log-src会收到类似如下一条日志：

```log
2020-04-15 10:53:21  level=info 收到日志 Hello
```

- 启动open-api

```shell
go run open-api/main.go
```


## MICRO API

```shell
micro api --handler=web --namespace=go.micro.learning.web
```

- 调用Sum-srv服务

```shell
http://localhost:8080/portal/sum?input=11
```

发现没有结果,需要修改文件portal-web/main.go

```go
service.HandleFunc("/learning/sum", Sum) 修改为
service.HandleFunc("/portal/sum", Sum)
```


```shell
micro api --handler=api --namespace=go.micro.learning.api
```

- 调用open-api服务

```shell
curl  "http://localhost:8080/open/Fetch?sum=12&prime=123"
```

结果如下：

```shell
{"prime":[0,1,2,3,5,7,11,13,17,19,23,29,31,37,41,43,47,53,59,61,67,71,73,79,83,89,97,101,103,107,109,113],"sum":78}
```

