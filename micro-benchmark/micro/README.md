# Micro 内部组件定制组合性能测试

我们在这一系列尝试给大家演示把Micro中的各种组件组合起来，测试它们之间各自配合的性能

## 运行环境

以下测试目前在单机上执行

- Go-Version: go version go1.12.6 darwin/amd64
- Go-Micro: v1.8.0
- 操作系统：MacOSX 10.14.5 (18F203)
- 内存：32G
- 处理器：2.2 GHz Intel Core i7

## Transport 对比

我们先对比目前Micro中提供有的transport，它们分别是Http（默认）、grpc、memery、quic、tcp（插件）、nats（插件）、rabbitmq（插件）、utp（插件）

### 对比结果

-|T+S|平均<br/>(ms)|中位<br/>(ms)|最大<br/>(ms)|最小<br/>(ms)|P90<br/>(ms)|P99<br/>(ms)|TPS
---|---|---|---|---|---|---|---|---
HttpTransport|100|148.780|113.004|802.520|0.437|302.865|396|664
TcpTransport|100|2.047|1.485|67.188|0.109|3.646|23|45310
gRpcTransport|100|3.158|2.572|38.030|0.190|5.201|20|30102
gRpc-TcpTransport|100|4.377|3.173|59.223|0.260|9.368|22|21734

### Http Transport

```bash
$ cd _default
$ go run server.go
# 切换窗口到同目录
$ go run client.go  -c 100 -n 100000

2019/08/11 21:47:31 concurrency: 100
requests per client: 1000

2019/08/11 21:47:31 message size: 677 bytes

2019/08/11 21:50:01 took 150601 ms for 100000 requests
2019/08/11 21:50:01 sent     requests    : 100000
2019/08/11 21:50:01 received requests    : 100000
2019/08/11 21:50:01 received requests_OK : 99966
2019/08/11 21:50:01 throughput  (TPS)    : 664
2019/08/11 21:50:01 concurrency mean    median  max     min     p90     p99     TPS
2019/08/11 21:50:01 100         148780257ns     113003500ns     802520000ns     437000ns        395687500ns     302865000ns     664
2019/08/11 21:50:01 100         148.780ms       113.004ms       802.520ms       0.437ms 302.865ms       396ms   664
```

### TCP Transport

```bash
$ cd default-tcp-transport
$ go run server.go
# 切换窗口到同目录
$ go run client.go  -c 100 -n 100000

2019/08/11 17:13:25 concurrency: 100
requests per client: 1000

2019/08/11 17:13:25 message size: 677 bytes

2019/08/11 17:13:27 took 2207 ms for 100000 requests
2019/08/11 17:13:27 sent     requests    : 100000
2019/08/11 17:13:27 received requests    : 100000
2019/08/11 17:13:27 received requests_OK : 100000
2019/08/11 17:13:27 throughput  (TPS)    : 45310
2019/08/11 17:13:27 concurrency	mean	median	max	min	p90	p99	TPS
2019/08/11 17:13:27 100 	2047433ns	1485000ns	67188000ns	109000ns	23027000ns	3646000ns	45310
2019/08/11 17:13:27 100 	2.047ms	1.485ms	67.188ms	0.109ms	3.646ms	23ms	45310
```

### grpc Transport

```bash
$ cd grpc
$ go run server.go
# 切换窗口到同目录
$ go run client.go  -c 100 -n 100000

2019/08/11 22:16:23 concurrency: 100
requests per client: 1000

2019/08/11 22:16:23 message size: 677 bytes

2019/08/11 22:16:26 took 3322 ms for 100000 requests
2019/08/11 22:16:26 sent     requests    : 100000
2019/08/11 22:16:26 received requests    : 100000
2019/08/11 22:16:26 received requests_OK : 100000
2019/08/11 22:16:26 throughput  (TPS)    : 30102
2019/08/11 22:16:26 concurrency mean    median  max     min     p90     p99     TPS
2019/08/11 22:16:26 100         3157656ns       2572000ns       38030000ns      190000ns        19897000ns      5201000ns       30102
2019/08/11 22:16:26 100         3.158ms 2.572ms 38.030ms        0.190ms 5.201ms 20ms    30102
```

### grpc-tcp Transport

```bash
$ cd grpc-tcp-transport
$ go run server.go
# 切换窗口到同目录
$ go run client.go  -c 100 -n 100000

2019/08/11 22:21:15 concurrency: 100
requests per client: 1000

2019/08/11 22:21:15 message size: 677 bytes

2019/08/11 22:21:19 took 4601 ms for 100000 requests
2019/08/11 22:21:19 sent     requests    : 100000
2019/08/11 22:21:19 received requests    : 100000
2019/08/11 22:21:19 received requests_OK : 100000
2019/08/11 22:21:19 throughput  (TPS)    : 21734
2019/08/11 22:21:19 concurrency mean    median  max     min     p90     p99     TPS
2019/08/11 22:21:19 100         4376550ns       3173000ns       59223000ns      260000ns        21572500ns      9368000ns       21734
2019/08/11 22:21:19 100         4.377ms 3.173ms 59.223ms        0.260ms 9.368ms 22ms    21734
```