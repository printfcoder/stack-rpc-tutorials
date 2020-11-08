# 多版本服务场景

- [v1服务及客户端](v1)
- [v2服务](v2)

## 启动

1. 运行V1, v2 与服务端

```bash
cd v1
go run srv.go &
cd ../v2
go run srv.go
```

2. 运行V1客户端

```bash
cd v1
go run cli.go
```

## 原理

Go-Micro 针对客户端选择服务时增加了选项WithSelectOption，可以通过声明它来选择使用哪些服务：

```go
func Filter(v string) client.CallOption {
	filter := func(services []*registry.Service) []*registry.Service {
		var filtered []*registry.Service

		for _, service := range services {
			if service.Version == v {
				filtered = append(filtered, service)
			}
		}

		return filtered
	}

	return client.WithSelectOption(selector.WithFilter(filter))
}
```

上面的过滤器中我们可以看到，**Filter**将版本与入参比较，取需要版本，我们也可以传入其它信息，进行其它大于或小于或更复杂情景的逻辑进行服务选择。