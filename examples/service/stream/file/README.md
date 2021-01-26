# File

本篇演示Stack中的大文件传输（流式调用）。

## 内容

- server.go - 文件接收服务端
- client.go - 文件发送客户端

## 运行File

使用protoc生成相应的代码
```
protoc --go_out=. --stack_out=. proto/file.proto
```

运行服务端

```shell
go run server.go
```

运行客户端

```shell
go run client.go
```

## 使用

使用Postman/CURL向[http://127.0.0.1:8080/upload](http://127.0.0.1:8080/upload)发送文件

```shell
curl --location --request POST 'http://127.0.0.1:8080/upload' \
--header 'Content-Type: multipart/form-data;' \
--form 'file=@"README.md"'
```

回参是该文件的md5