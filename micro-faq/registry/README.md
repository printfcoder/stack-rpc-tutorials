# 注册常见问题

1. 多网卡时注册使用了不需要的IP地址

答：可以使用`-server_address=ip:port` 选项启动服务，如：`go run main.go -server_address=127.0.0.1:8876`，这样就可以指定服务的IP地址。