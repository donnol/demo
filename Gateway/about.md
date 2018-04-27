# Gateway 网关

网关，会承载巨大的流量，那么如何保证整个系统的稳定性和高可用，保证高性能和可靠，以支撑业务呢？

## [京东网关架构](https://my.oschina.net/u/3772106/blog/1799394)

    1 统一接入
    2 引入 NIO
    3 分离请求解析和业务处理
    4 降级
    5 限流
    6 熔断
    7 快速失败
    8 监控统计