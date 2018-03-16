# 分布式系统

[一组电脑（computer），通过网络相互链接传递消息与通信并协调它们的行为而形成的系统。把需要进行大量计算的工程数据分区成小块，由多台计算机分别计算，再上传运算结果后，将结果统一合并得出数据结论的科学](https://zh.wikipedia.org/wiki/%E5%88%86%E5%B8%83%E5%BC%8F%E8%AE%A1%E7%AE%97)

常见分布式系统：面向服务(SOA)，大型多人在线游戏(MMOG)，对等网络(P2P)

1 [CAP 定理](https://zh.wikipedia.org/wiki/CAP%E5%AE%9A%E7%90%86)

    C(Consistence): 一致性
        所有节点访问同一份最新的数据副本
    A(Availability): 可用性
        每次请求都能获取到非错的响应，但是不保证获取的数据为最新数据
    P(Network partitioning): 分区容忍性
        分区相当于对通信的时限要求。系统如果不能在时限内达成数据一致性，就意味着发生了分区的情况

分布式系统只能满足三项中的两项而不可能满足全部三项

相关

    ACID: 传统事务模型，追求强一致性
    BASE: 基本可用（Basically Available）、软状态（Soft State）、最终一致性（Eventual Consistency）
