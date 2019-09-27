# gRPC FAQ

1. 纯gRPC服务能与micro应用互调吗？

答：不可以均衡互调，只能直连互调，micro是micro，gRPC是gRPC。micro可以直连调用gRPC，与其它语言一致，需要自行配置彼此服务地址。gRPC可以直连调用`grpc.Newservice`服务。Micro有自己的风格与协议，与gRPC是不同的，Micro通过Client组件均衡调用其它服务（Server）的Transport组件完成rpc请求，gRPC中是没有这两个组件的，也就无法均衡与直接接收请求。
更多示例，请参考：[gRPC示例](https://github.com/micro-in-cn/tutorials/tree/master/examples/middle-practices/micro-grpc)