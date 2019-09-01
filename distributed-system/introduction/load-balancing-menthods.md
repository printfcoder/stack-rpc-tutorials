# 负载均衡方法

负载均衡是一种将N个请求分发向M个服务的算法或机制，它的目的是有效地将请求与流量颁发到在服务注册池中的服务器。负载均衡的目地最大化增强资源的利用率，以最小的时间得到最大的产出。

## 常见的负载均衡方法

- Round robin **轮询**，以顺序的方式将N个请求轮流发到M个服务上，是最简单最轻量的方法，[参考](https://en.wikipedia.org/wiki/Round-robin_DNS)
- Weighted round robin **权重化轮询**，在**轮询**的基础上为每个服务器分配不同的权重，权重高的接收更多的请求。[参考](https://en.wikipedia.org/wiki/Weighted_round_robin)
- Round-robin DNS **DNS轮询**，M个服务的器的IP地址绑到一个域名上，客户端会以轮询的方式得到一个带过期时间的服务端IP，一定时间内会使用这个ip进行请求，过期会颁发新的服务器ip。
- Least Connection **最小连接**， 为了减小高负载的服务器压力，该方法将请求分发给当前压力比较小的服务器。[参考](https://docs.citrix.com/en-us/netscaler/12/load-balancing/load-balancing-customizing-algorithms/leastconnection-method.html)
- Least Response Time **最小响应时间**，以服务器的*连接数*和*平均响应时间*作为标准，将请求转给最合适的服务器，[参考](https://docs.citrix.com/en-us/netscaler/12/load-balancing/load-balancing-customizing-algorithms/leastresponsetime-method.html)
- Weighted least connection **权重化最小连接**，在**最小连接**的方法基础上再根据服务器的运算能力分配相应的权重，按权重与最小连接数以一定算法选取服务下发请求。
- Weighted response **响应权重**，基于服务器从接收请求到响应所花的时间采样计算其动态权重，权重高分到请求概率高。
- Hash **哈希**，将客户端IP与服务端IP端口（或其它唯一信息）采用一定算法混合哈希得到一个唯一值后，通过这个哈希值每次都到对应的服务器中。[参考](https://docs.citrix.com/en-us/netscaler/12/load-balancing/load-balancing-customizing-algorithms/hashing-methods.html)
