# 网络协议

## [开放式系统互联通信参考模型(Open System Interconnection Reference Model，缩写为 OSI)](https://zh.wikipedia.org/wiki/OSI%E6%A8%A1%E5%9E%8B)

    7 应用层 - 提供为应用软件而设的界面，如 HTTP，HTTPS，FTP，TELNET，SSH，SMTP，POP3
    6 表示层 - 把数据转换为能与接收者的系统格式兼容并适合传输的格式
    5 会话层 - 在数据传输中设置和维护电脑网络中两台电脑之间的通信连接
    4 传输层 - 把传输表头（TH）加至数据以形成数据包，如 TCP, UDP
    3 网络层 - 决定数据的路径选择和转寄，将网络表头（NH）加至数据包，以形成分组，如 IP
    2 数据链路层 - 网络寻址、错误侦测和改错，如以太网、无线局域网（Wi-Fi）和通用分组无线服务（GPRS）
    1 物理层 - 在局域网上传送帧，它负责管理电脑通信设备和网络媒体之间的互通，如针脚、电压、线缆规范、集线器、中继器、网卡、主机适配器

## Socket

[使用](https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/08.1.md)
