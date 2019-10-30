# micro使用etcd作为配置中心

<a name="zTEpU"></a>
## 概述

根据micro所有组件可插拔的特性，只要实现以下接口，即可注册成为micro项目的配置中心。<br />本文涉及到的etcd配置中心组件，即是官方根据以下接口实现，详情查看[源码](https://github.com/micro/go-micro/tree/master/config/source/etcd)。

```go
// Source is the source from which config is loaded
type Source interface {
	Read() (*ChangeSet, error)
	Watch() (Watcher, error)
	String() string
}
```


<a name="yMayI"></a>
## 快速搭建etcd实验环境

docker快速启动<br />访问地址为：         你的主机ip:2379<br />或者本地访问：   127.0.0.1:2379

```shell
docker network create etcd_net && docker run -d --name etcd-micro \
    --restart always   \
    --network etcd_net \
    --publish 2379:2379 \
    --publish 2380:2380 \
    --env ALLOW_NONE_AUTHENTICATION=yes \
    --env ETCD_ADVERTISE_CLIENT_URLS=http://etcd-micro:2379 \
    bitnami/etcd:latest
```

进入容器中

```shell
docker exec -it etcd-micro /bin/bash
```

使用etcdctl设置键值，要注意符合json格式

```shell
etcdctl put /micro/app01/config '{"name":"etcd_micro","env":"dev","address":"127.0.0.1","port":3307}'
```


<a name="wC6wZ"></a>
## 设置etcd配置中心

定义项目需要用到的配置

```go
type Appconfig struct {
	Name    string `json:"name"`
	Env     string `json:"env"`
	Address string `json:"address"`
	Port    int64  `json:"port"`
}
```

 初始化 etcd配置中心source

```go
etcdSource := etcd.NewSource(

		etcd.WithAddress("192.168.172.3:2379"),

		// 限定键的前缀，默认是 /micro/config
		etcd.WithPrefix("/micro/app01"),

		// optionally strip the provided prefix from the keys, defaults to false
		//etcd.StripPrefix(false),

	)
```

将该source即配置源设置到micro的config中，这里有两种设置方法。

```go
	// 1
    conf1 := config.NewConfig(config.WithSource(etcdSource))

    // 2
    conf2 := config.NewConfig()
    err:=conf2.Load(etcdSource)
    if err != nil {
        log.Fatal(err)
    }
```

使用键从config中获取到配置，设置到Appconfig实例中

```go
    var Myconf Appconfig
    Myconf = Appconfig{}
    err := conf.Get("micro", "app01", "config").Scan(&Myconf)
	
	if err != nil {
		log.Fatal(err)
	}
```
 

