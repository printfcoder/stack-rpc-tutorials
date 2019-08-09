# Micro 内部组件定制组合性能测试

## 运行环境

Go-Micro: v1.8.0

### 单机

操作系统：MacOSX 10.14.5 (18F203)
内存：32G
处理器：2.2 GHz Intel Core i7

## HttpTransport

```bash
$ go run client.go -c 100 -n 100000
2019/08/09 11:15:20 Servers: [127.0.0.1:8972]

2019/08/09 11:15:20 concurrency: 100
requests per client: 1000

2019/08/09 11:15:20 message size: 581 bytes

2019/08/09 11:17:45 took 145219 ms for 100000 requests
2019/08/09 11:17:45 sent     requests    : 100000
2019/08/09 11:17:45 received requests    : 100000
2019/08/09 11:17:45 received requests_OK : 0
2019/08/09 11:17:45 throughput  (TPS)    : 688
2019/08/09 11:41:23 concurrency mean    median  max     min     p90     p99     TPS
2019/08/09 11:41:23 10 :        143288995ns     135104500ns     353413000ns     100120000ns     106435500ns
2019/08/09 11:17:45 mean: 143288995 ns, median: 135104500 ns, max: 353413000 ns, min: 100120000 ns, p99: 271410000 ns
2019/08/09 11:17:45 mean: 143 ms, median: 135 ms, max: 353 ms, min: 100 ms, p99: 271 ms
```

## TCP Transport