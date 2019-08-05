### 操作步骤
1. 安装nsq，分别获取nsqlookupd & nsqd节点列表地址
2. 替换```config/config.go```中```nsqLookupdAddrs```和```nsqdAddrs```中的值
3. 启动client: ```go run client.go```
4. 启动server: ```go run server.go```
