#使用metadata从客户端传递非业务数据给服务端


客户端用JSON客户端的代码
- [JSON客户端](../json)


我们先来看看代码
```go
func main() {
	cli := client.NewClient(
		// 与目录服务同注册中心即可
		client.Registry(mdns.NewRegistry()),
	)

	// 调用目标服务的结构
	req := cli.NewRequest("go.micro.rpc.example", "Example.Call",
		&whatEverReq{
			Name: "John",
		},
		// 不确定对方服务时，需要使用JSON格式，而不是protobuf
		client.WithContentType("application/json"))

	// 自定义元数据
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id": "john",
		"X-From-Id": "script",
	})

	rsp := &whatEverRsp{}

	// 调用服务
	if err := cli.Call(ctx, req, rsp); err != nil {
		fmt.Println("call err: ", err, rsp)
		return
	}

	fmt.Println("rsp: ", rsp.Message)
}

```

大家注意了 `自定义元数据` 就是这段代码


```go
    // 自定义元数据
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id": "john",
		"X-From-Id": "script",
	})

```

在选择key的时候一定要注意，micro服务端接收到的元数据的key是符合http头消息的命名规则的
假设在客户端你选择的key 是 "x-user-ID" 服务端接收到的key 仍然是 "X-User-Id"。

###为了在编码过程中不产生困扰，元数据的key 建议遵从http头消息的命名规则(单词首字母大写，以中划线划分单词)。


服务端代码我们就用

- [rpc服务端](../rpc)

的代码，我们对Call函数做如下修改
```go
func (e *Example) Call(ctx context.Context, req *proto.CallRequest, rsp *proto.CallResponse) error {
	log.Printf("收到 Example.Call 请求 %v\n", req)
	fmt.Printf("%v\n", req)

+	if md, ok := metadata.FromContext(ctx); ok {
+		log.Printf("metadata X-User-Id=%s", md["X-User-Id"])
+	}

	if len(req.Name) == 0 {
		return errors.BadRequest("go.micro.rpc.example", "no content")
	}

	rsp.Message = "RPC Call收到了你的请求 " + req.Name
	return nil
}

```
追加了3行代码
```go
    if md, ok := metadata.FromContext(ctx); ok {
        log.Printf("metadata X-User-Id=%s", md["X-User-Id"])
    }
```

修改完成后，启动rpc服务端，再运行json客户端。
rpc端的log如下

```bash
name:"John"
2019/10/18 12:56:25 metadata X-User-Id=john
```
