# 同步请求相关FAQ

## 1. 多版本micro server同时存在，micro client是否会根据version调取相对应的server？

答，不会。Micro不会默认选择，这是因为版本号是自定义的，没有可比较的默认方案。不过Go-micro支持客户端自行选择需要服务端，可以通过**client.WithSelectOption**自定义，参考：[version-filter](https://github.com/micro-in-cn/tutorials/tree/master/examples/senior-practices/micro-filter/version)。

## 2. 客户端调用服务端出现408错误问题？

答，请按照以下思路检测：
  a) 注册中心是否有未卸载干净的同名幽灵服务

  b) 确认两个网络能彼此通信，未禁ping的，使用ping命令即可

  c) 服务是否有多个网卡，比如虚拟网卡，go-micro会使用它拿到的第一个网卡的ip地址，此时可能会拿到其它机器无法访问的虚拟网卡的ip，如果有多个，请使用能通信上的ip地址，并使用server_address指令启动，如`go run main.go --server_address=192.168.3.5:8080`
