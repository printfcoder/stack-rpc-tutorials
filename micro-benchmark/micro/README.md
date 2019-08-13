# Micro 内部组件定制组合性能测试

我们在这一系列尝试给大家演示把Micro中的各种组件组合起来，测试它们之间各自配合的性能

## 运行环境

以下测试目前在单机上执行

- Go-Version: go version go1.12.6 darwin/amd64
- Go-Micro: v1.8.3
- 操作系统：MacOSX 10.14.5 (18F203)
- 内存：32G
- 处理器：2.2 GHz Intel Core i7

## Transport 对比

我们先对比目前Micro中提供有的transport，它们分别是Http（默认）、grpc、memory、quic、tcp（插件）、nats（插件）、rabbitmq（插件）、utp（插件）

其中quic因为其go包经常变动，修改麻烦，暂不提供测试。未来Micro会改成基于更加稳定的kcp协议，二者都是基于UDP开发的协议。

### 对比结果

-|平均<br/>(ms)|中位<br/>(ms)|最大<br/>(ms)|最小<br/>(ms)|P90<br/>(ms)|P99<br/>(ms)|TPS
---|---|---|---|---|---|---|---
HttpTransport（1）|148.780|113.004|802.520|0.437|302.865|396|664
TcpTransport|2.047|1.485|67.188|0.109|3.646|23|45310
GRPCTransport|3.158|2.572|38.030|0.190|5.201|20|30102
GRPC-TCPTransport|3.064|2.562|36.174|0.214|4.871|19|30911
UTPTransport|6.757|6.583|41.935|0.136|8.402|24|14388
NATsTransport（2）|20.113|18.007|1003.489|1.697|28.260|596|4865
NATsTransport（3）|14.364|13.280|80.938|1.474|21.289|59|6773
RabbitMQTransport（4）|127.019|22.259|5005.101|1.866|32.702|5004|700

> * 1 http模式与具体服务器对Http请求的出入限制配置有关，待优化配置再测试
> * 2 nats运行在本机Docker中，分配内存为2GB，CPU核心数为6核
> * 3 nats运行在本机Docker中，分配内存为4GB，CPU核心数为6核，当适当增加内存为双倍时，性能有40%左右提升，但与tpc，grpc仍有不小差距。
> * 4 rabbit运行在本机Docker中，分配内存为4GB，CPU核心数为6核，性能极差，需要调优再测试

上面的表格中可见，TCP>grpc≈grpc-tcp>utp>nats>http>rabbitMQ

具体数值在不同宿主机中有差异，视具体情况而定，但是总体趋势是有参考意义的。

### Http Transport

[Http-Transport](./http-transport)

```bash
$ cd http-transport
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

[TCP-Transport](./tcp-transport)

```bash
$ cd tcp-transport
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

### GRPC Transport

[GRPC-Transport](./grpc-transport)

```bash
$ cd grpc-transport
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

### GRPC-TCP Transport

[GRPC-TCP-Transport](./grpc-tcp-transport)

```bash
$ cd grpc-tcp-transport
$ go run server.go
# 切换窗口到同目录
$ go run client.go  -c 100 -n 100000

2019/08/11 22:21:20 concurrency: 100
requests per client: 1000

2019/08/11 22:21:20 message size: 677 bytes

2019/08/12 22:21:20 sent     requests    : 100000
2019/08/12 22:21:20 received requests    : 100000
2019/08/12 22:21:20 received requests_OK : 100000
2019/08/12 22:21:20 throughput  (TPS)    : 30911
2019/08/12 22:21:20 concurrency mean    median  max     min     p90     p99     TPS
2019/08/12 22:21:20 100         3064103ns       2562000ns       36174000ns      214000ns        18531500ns      4871000ns       30911
2019/08/12 22:21:20 100         3.064ms 2.562ms 36.174ms        0.214ms 4.871ms 19ms    30911
```

### UTP Transport

[UTP-Transport](./utp-transport)

```bash
$ cd utp-transport
$ go run server.go
# 切换窗口到同目录
$ go run client.go  -c 100 -n 100000

2019/08/11 23:33:58 concurrency: 100
requests per client: 1000

2019/08/11 23:33:58 message size: 677 bytes

2019/08/11 23:34:05 took 6950 ms for 100000 requests
2019/08/11 23:34:05 sent     requests    : 100000
2019/08/11 23:34:05 received requests    : 100000
2019/08/11 23:34:05 received requests_OK : 100000
2019/08/11 23:34:05 throughput  (TPS)    : 14388
2019/08/11 23:34:05 concurrency mean    median  max     min     p90     p99     TPS
2019/08/11 23:34:05 100         6757423ns       6583000ns       41935000ns      136000ns        24452000ns      8402000ns       14388
2019/08/11 23:34:05 100         6.757ms 6.583ms 41.935ms        0.136ms 8.402ms 24ms    14388
```

### NATs Transport

[NATs-transport](./nats-transport)

```bash
$ cd nats-transport
$ go run server.go
# 切换窗口到同目录
$ go run client.go  -c 100 -n 100000

2019/08/11 00:01:08 concurrency: 100
requests per client: 1000

2019/08/14 00:11:56 message size: 677 bytes

2019/08/14 00:12:17 took 20552 ms for 100000 requests
2019/08/14 00:12:17 sent     requests    : 100000
2019/08/14 00:12:17 received requests    : 100000
2019/08/14 00:12:17 received requests_OK : 100000
2019/08/14 00:12:17 throughput  (TPS)    : 4865
2019/08/14 00:12:17 concurrency mean    median  max     min     p90     p99     TPS
2019/08/14 00:12:17 100         20112930ns      18007000ns      1003489000ns    1697000ns       595746000ns     28260000ns      4865
2019/08/14 00:12:17 100         20.113ms        18.007ms        1003.489ms      1.697ms 28.260ms        596ms   4865
```

### RabbitMQ Transport

[RabbitMQ-transport](./rabbitmq-transport)

```bash
2019/08/14 00:29:22 concurrency: 100
requests per client: 1000

2019/08/14 00:29:22 message size: 677 bytes

2019/08/14 00:42:20 sent     requests    : 100000
2019/08/14 00:42:20 received requests    : 100000
2019/08/14 00:42:20 received requests_OK : 97908
2019/08/14 00:42:20 throughput  (TPS)    : 700
2019/08/14 00:42:20 concurrency mean    median  max     min     p90     p99     TPS
2019/08/14 00:42:20 100         127019375ns     22259000ns      5005101000ns    1866000ns       5003709000ns    32702000ns      700
2019/08/14 00:42:20 100         127.019ms       22.259ms        5005.101ms      1.866ms 32.702ms        5004ms  700
```
