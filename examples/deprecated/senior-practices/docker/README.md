# Docker

很多朋友有疑问为什么在Docker中运行时会调不通服务，这是因为Docker有自己的网络，我们的服务端与客户端都要在这个网络中，才可能彼此访问。下面我们就给大家演示纯Docker环境时，怎么发布的简单流程。


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

5. 同样的套路编译打包client

```bash
$ GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o go-micro-demo-client client.go
$ docker build . -t go-micro-demo-client:latest
```

6. 启动客户端容器，输出结果

```bash
$ docker run go-micro-demo-client
你好，Micro中国
```
