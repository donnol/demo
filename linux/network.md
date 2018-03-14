# 网络

1 网络地址

[唯一本地地址(ULA)](https://linux.cn/article-9434-1.html)

    使用 fd00::/8 地址块
    类似于我们常用的 IPv4 的私有地址
        IPv4 的私有地址分类和网络地址转换（NAT）功能是为了缓解 IPv4 地址短缺的问题
    与本地链路地址link-local addresses（fe80::/10）的区别
        一是，本地链路地址是不可路由的，因此，你不能跨子网使用它
        二是，ULA 是你自己管理的；你可以自己选择它用于子网的地址范围，并且它们是可路由的
    ULA 是仅为私有网络使用的，并且应该阻止其流出你的网络，不允许进入因特网。这很简单，在你的边界设备上只要阻止整个 fd00::/8 范围的 IPv6 地址即可实现
    配置
        radvd 路由器公告守护程序
    ULA 转发 IPv6后，就可以使用一个合适的主机名来代替长长的 IPv6 地址

2 地址解析协议(ARP, address resolution protocol)

[通过解析网路层地址来找寻数据链路层地址的一个在网络协议包中极其重要的网络传输协议](https://zh.wikipedia.org/wiki/%E5%9C%B0%E5%9D%80%E8%A7%A3%E6%9E%90%E5%8D%8F%E8%AE%AE)

    1.当发送数据时，主机A会在自己的ARP缓存表中寻找是否有目标IP地址。如果找到就知道目标MAC地址为（00-BB-00-62-C2-02），直接把目标MAC地址写入帧里面发送就可。

    2.如果在ARP缓存表中没有找到相对应的IP地址，主机A就会在网络上发送一个广播（ARP request），目标MAC地址是“FF.FF.FF.FF.FF.FF”，这表示向同一网段内的所有主机发出这样的询问：“192.168.38.11的MAC地址是什么？”

    3.网络上其他主机并不响应ARP询问，只有主机B接收到这个帧时，才向主机A做出这样的回应（ARP response）：“192.168.38.11的MAC地址是00-BB-00-62-C2-02”，此回应以单播方式。这样，主机A就知道主机B的MAC地址，它就可以向主机B发送信息。同时它还更新自己的ARP高速缓存（ARP cache），下次再向主机B发送信息时，直接从ARP缓存表里查找就可。

[ARP 欺骗](http://www.cnblogs.com/manue1/p/4462327.html)

    ARP协议是建立在信任局域网内所有节点的基础上的，它的效率很高。但是不安全。它是无状态的协议。它不会检查自己是否发过请求包，也不知道自己是否发过请求包。它也不管是否合法的应答，只要收到目标mac地址是自己的ARP reply或者ARP广播包（包括ARP reply和ARP request），都会接受并缓存

    监控工具
        arp
        nbtscan
    防御
        双向绑定
        防火墙
        DHCP服务器
        划分安全区域

3 [DHCP 动态主机设置协议](https://zh.wikipedia.org/wiki/%E5%8A%A8%E6%80%81%E4%B8%BB%E6%9C%BA%E8%AE%BE%E7%BD%AE%E5%8D%8F%E8%AE%AE)

    一个局域网的网络协议，使用UDP协议工作
        用于内部网络或网络服务供应商自动分配IP地址给用户
        用于内部网络管理员作为对所有电脑作中央管理的手段
    使网络管理员能够集中管理和自动分配IP网络地址的通信协议

[与 NAT 区别](https://www.zhihu.com/question/66893227)

    DHCP 分配局域网内部IP
    NAT 将局域网IP转换为公网IP