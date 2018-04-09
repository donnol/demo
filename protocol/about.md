# 网络协议

## [开放式系统互联通信参考模型(Open System Interconnection Reference Model，缩写为 OSI)](https://zh.wikipedia.org/wiki/OSI%E6%A8%A1%E5%9E%8B)

    7 应用层 - 提供为应用软件而设的界面，如 HTTP，HTTPS，FTP，TELNET，SSH，SMTP，POP3
    6 表示层 - 把数据转换为能与接收者的系统格式兼容并适合传输的格式
    5 会话层 - 在数据传输中设置和维护电脑网络中两台电脑之间的通信连接
    4 传输层 - 把传输表头（TH）加至数据以形成数据包，如 TCP, UDP
    3 网络层 - 决定数据的路径选择和转寄，将网络表头（NH）加至数据包，以形成分组，如 IP
    2 数据链路层 - 网络寻址、错误侦测和改错，如以太网、无线局域网（Wi-Fi）和通用分组无线服务（GPRS），如 MAC
    1 物理层 - 在局域网上传送帧，它负责管理电脑通信设备和网络媒体之间的互通，如针脚、电压、线缆规范、集线器、中继器、网卡、主机适配器

## 其它

[ARP-Address Resolution Protocol: 通过解析网络层地址(3)来找寻数据链路层地址(2)的网络传输协议](https://zh.wikipedia.org/wiki/%E5%9C%B0%E5%9D%80%E8%A7%A3%E6%9E%90%E5%8D%8F%E8%AE%AE)

    在网路层和数据链接层之间实现
    在 IPv6 中使用 NDP(邻居发现协议) 代替 ARP

[CDP-Cisco Discovery Protocol: 思科发现协议，运行在 Cisco IOS 的路由器和交换机之间](https://en.wikipedia.org/wiki/Cisco_Discovery_Protocol)

    在数据链路层实现

[DHCP-Dynamic Host Configuration Protocol: 动态主机设置协议，是一个局域网的网络协议，使用 UDP 协议工作](https://zh.wikipedia.org/wiki/%E5%8A%A8%E6%80%81%E4%B8%BB%E6%9C%BA%E8%AE%BE%E7%BD%AE%E5%8D%8F%E8%AE%AE)

    用于内部网络或网络服务供应商自动分配IP地址给用户
    用于内部网络管理员作为对所有电脑作中央管理的手段

[NAT-Netword Address Translation: 网络地址转换，一种在 IP 数据包通过路由器或防火墙时重写来源 IP 地址或目的 IP 地址的技术](https://zh.wikipedia.org/wiki/%E7%BD%91%E7%BB%9C%E5%9C%B0%E5%9D%80%E8%BD%AC%E6%8D%A2)

    一种解决IPv4地址短缺以避免保留IP地址困难的方案
        IPv6的广泛采用将使得NAT不再需要
    专有IP转为公有IP
    普遍使用在有多台主机但只通过一个公有IP地址访问因特网的私有网络中
    NAT也让主机之间的通信变得复杂，导致了通信效率的降低

[DNS-Domain Name System: 域名系统，是互联网的一项服务。它作为将域名和 IP 地址相互映射的一个分布式数据库，能够使人更方便地访问互联网](https://zh.wikipedia.org/wiki/%E5%9F%9F%E5%90%8D%E7%B3%BB%E7%BB%9F)

    使用TCP和UDP端口53
    www.example.com = www.example.com.root = www.example.com.
        根域名.root对于所有域名都是一样的，所以平时是省略的

[DNS 加密](https://www.v2ex.com/t/445199#reply10)

    防止 ISP 对 DNS 协议的恶意劫持，或者部分 DNS 服务对 NXDOMAIN 的劫持（对解析失败的域名跳转到广告页面
    工具：DNSCrypt-Proxy

[DNS over HTTPS](https://www.v2ex.com/t/365828)

[ICMP-Internet Control Message Protocol：互联网控制消息协议，是互联网协议族的核心协议之一](https://zh.wikipedia.org/wiki/%E4%BA%92%E8%81%94%E7%BD%91%E6%8E%A7%E5%88%B6%E6%B6%88%E6%81%AF%E5%8D%8F%E8%AE%AE)

    用于TCP/IP网络中发送控制消息，提供可能发生在通信环境中的各种问题反馈，通过这些信息，使管理者可以对所发生的问题作出诊断，然后采取适当的措施解决
    例子之一是TTL值过期
        每个路由器在转发数据报的时候都会把IP包头中的TTL值减1
        如果TTL值为0，“TTL在传输中过期”的消息将会回报给源地址
        每个ICMP消息都是直接封装在一个IP数据包中的，因此，和UDP一样，ICMP是不可靠的
    ICMP报头从IP报头的第160位开始（IP首部20字节）（除非使用了IP报头的可选部分）

[PPPoE-Point-to-Point Protocol Over Ethernet: 以太网上的点对点协议，是将点对点协议（PPP）封装在以太网（Ethernet）框架中的一种网络隧道协议](https://zh.wikipedia.org/wiki/PPPoE)

    由于协议中集成PPP协议，所以实现出传统以太网不能提供的身份验证、加密以及压缩等功能，也可用于缆线调制解调器（cable modem）和数位用户线路（DSL）等以以太网协议向用户提供接入服务的协议体系

[USB-Universal Serial Bus: 通用串行总线，是连接计算机系统与外部设备的一种串口总线标准，也是一种输入输出接口的技术规范](https://zh.wikipedia.org/wiki/USB)

    最大的特点是支持热插拔和即插即用
    在速度上远比并行端口（例如EPP、LPT）与串行接口（例如RS-232）等传统电脑用标准总线快上许多

## Socket

[使用](https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/08.1.md)

## 抓包

    工具：Wireshark/tshark, Fiddler, Driftnet, HttpWatch, Charles, tcpdump, gopacket, pproxy
    库: libpcap
        winpcap

## [路由器](https://zh.wikipedia.org/wiki/%E8%B7%AF%E7%94%B1%E5%99%A8)

    工作在OSI模型的第三层——即网络层

    路由器与交换机在概念上有一定重叠但也有不同：
        交换机泛指工作于任何网络层次的数据中继设备（尽管多指网桥），而路由器则更专注于网络层

    From zhihu
        交换机工作于数据链路层，用来隔离冲突域，连接的所有设备同属于一个广播域（子网），负责子网内部通信
        路由器工作于网络层，用来隔离广播域（子网），连接的设备分属不同子网，工作范围是多个子网之间，负责网络与网络之间通信
