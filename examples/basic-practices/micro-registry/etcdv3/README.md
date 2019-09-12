# etcdv3

本篇演示如何使用etcdv3。

## 内容

- server.go - 服务端
- client.go - 客户端
- plugins.go - etcdv3插件

## 运行

运行服务端

```shell
go run server.go plugins.go --registry=etcdv3
```

运行客户端

```shell
go run client.go plugins.go --registry=etcdv3
```

## 其它

**Go-micro**中的原生注册包中是没有etcd的，故而`micro api`中也是没有的，因为micro指令是基于Go-Micro开发的。那怎么才能使指令`micro api`能够注册到etcd上呢？

我们就以etcdv3为例，给大家演示如何使用。

1. 下载micro源码

```bash
git clone git@github.com:micro/micro.git
```

或者

```bash
git clone https://github.com/micro/micro.git
```

或者

```bash
go get -u github.com/micro/micro
cd $GOPATH/src/github.com/micro/micro
```

2. 编译

2.1 切到源码目录

```bash
cd micro源码目录
```

2.2 增加插件文件plugins.go:

增加plugins.go到micro包的根目录

plugins.go

```golang
package main

import (
_ "github.com/micro/go-plugins/registry/etcdv3"
)
```

2.3 编译

```bash
go build -o mainWithEtcdv3 main.go plugins.go 
```

3. 运行

```bash
./mainWithEtcdv3 --registry=etcdv3 api
```
