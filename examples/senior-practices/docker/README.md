# Docker

1. 编译Server

```bash
$ GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o go-micro-demo-server server.go
```

2. 查看Consul容器的IP地址

我们假设还不知道Consul在Docker网络中的地址，现在我们获取一下地址：

```bash
$ docker inspect 797e51f0f699 | grep IPAddress
## 我们简单过滤一下，不考虑格式
            "SecondaryIPAddresses": null,
            "IPAddress": "172.17.0.3",
                    "IPAddress": "172.17.0.3",

```

其中的`797e51f0f699`是Consul容器的id

3. 打包docker镜像

我们需要将上一步中的consul容器的ip拿到放到Dockerfile中去

```Dockerfile
FROM alpine

ENV MICRO_REGISTRY consul
ENV MICRO_REGISTRY_ADDRESS 172.17.0.3:8500

RUN apk update && apk add tzdata && cp -r -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

ADD go-micro-demo-server /go-micro-demo-server

WORKDIR /
ENTRYPOINT [ "/go-micro-demo-server" ]
```

打包

```bash
$ docker build . -t go-micro-demo-server:latest
```

4. 启动Server容器

```bash
$ docker run go-micro-demo-server
```

5. 同样的套路编译打包client，并启动

```bash
$ GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o go-micro-demo-client client.go
$ docker build . -t go-micro-demo-client:latest
$ docker run go-micro-demo-client
```


得到容器的ip
