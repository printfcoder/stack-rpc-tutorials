# 同步请求相关FAQ

1. 多版本micro server同时存在，micro client是否会根据version调取相对应的server？

答，不会。Micro不会默认选择，这是因为版本号是自定义的，没有可比较的默认方案。不过Go-micro支持客户端自行选择需要服务端，可以通过**client.WithSelectOption**自定义，参考：[version-filter](https://github.com/micro-in-cn/tutorials/tree/master/examples/senior-practices/micro-filter/version)。