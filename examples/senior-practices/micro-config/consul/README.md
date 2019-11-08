# micro使用consul作为配置中心
#概述

micro-config的consul插件从consul的key/values读取配置

#准备工作

- 安装consul[安装](https://learn.hashicorp.com/consul/getting-started/install.html)

#数据存储格式

consul插件默认的key的前缀是/micro/config，这个是可以自定义的
我们在consul中存储数据
执行命令：
```
    consul kv put micro/app/mysql '{"address":"127.0.0.1","port":3306}'
```
#设置配置中心

指定配置来源
```go
    consulSource := consul.NewSource(
    	consul.WithAddress("127.0.0.1:8500"),
    	consul.WithPrefix("/micro/app"),
    )
```

下载配置源

```go
    // 创建配置实例
    conf := config.NewConfig()
    
    // 加载consul配置中心配置
    conf.Load(consulSource)
```

获取配置
```go
    var Conf Appconfig
    Conf = Appconfig{}

    err = conf.Get("micro", "app", "mysql").Scan(&Conf)
    if err != nil {
        log.Fatal(err)
    }
```

监听配置变化
```go
    watcher,err := conf.Watch()
    if err != nil {
        log.Fatal(err)
    }

    log.Logf("Watch changes ...")
    for {
        v, err := watcher.Next()
        if err != nil {
        log.Fatal(err)
    }

        log.Logf("Watch changes: %v", string(v.Bytes()))
    }
```
