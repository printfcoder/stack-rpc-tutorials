# gRPC 资源服务

## 服务端

```bash
cd config-srv
go run main.go
```

## 客户端

```bash
go run client/main.go
```

### 改动文件

切到config-srv/conf目录，改动**micro.yml**文件的任意值。

比如

```bash
micro:
  name: Micro
  version: 1.0.0
  hi: hello
```

改成

```bash
micro:
  name: Micro
  version: 1.0.0
  hi: I am fine
```

客户端会打印：

```bash
2019/04/28 10:57:15 Watch changes: {"hi":"I am fine","name":"Micro","version":"1.0.0"}
``` 