## 使用方式

使用protoc生成相应的代码

go

```bash
protoc --proto_path=$GOPATH:. --go_out=./api/proto --micro_out=./api/proto api.proto
```

python

```bash
python -m grpc_tools.protoc --proto_path=:.  --python_out=./python/  --grpc_python_out=./python/ api.proto
```

>  会在proto_path路径下搜索proto/api.proto中导入proto文件

再运行本api服务
--endpoint=go.micro.proxy.greeter
```
go run api.go
```

