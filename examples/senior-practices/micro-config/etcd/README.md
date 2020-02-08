# micro使用etcd作为配置中心

本示例我们演示使用ETCD作为配置中心

## 预置条件

没有安装可参考[官网教程](http://play.etcd.io/)安装。

因为我们要在`/micro/app`下创建配置，所以我们先执行以下命令，创建一个虚拟的Mysql配置

```bash
etcdctl put /micro/app/config "{\"micro\":{\"config\":{\"mysql\":{\"port\":3309}}}}"

OK
```

查看是否添加成功

```bash
etcdctl get /micro/app/config

/micro/app/config
{micro:{config:{mysql:{port:3306}}}}
```

## 设置etcd配置中心

初始化etcd配置中心source，见[源码](./main.go)

```go
	etcdSource := etcd.NewSource(
		etcd.WithAddress("127.0.0.1:2379"),
		etcd.WithPrefix("/micro/app"),
	)
```

将该source即配置源设置到micro的config中。

```go
	conf := config.NewConfig()
	err := conf.Load(etcdSource)
	if err != nil {
		log.Fatal(err)
	}
```

获取配置，设置到microCfg实例中

```go
	microCfg := &struct {
		Micro Micro `json:"micro"`
	}{}

	v := conf.Get("micro", "app", "config")
	err = v.Scan(&microCfg)
	if err != nil {
		log.Fatal(err)
	}

	log.Logf("Read mysql port: %d", microCfg.Micro.Config.Mysql.Port)
```
 

