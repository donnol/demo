# tls(Transport Layer Security) 传输层安全性协议

[前身是安全套接层（Secure Sockets Layer，缩写作 SSL），目的是为互联网通信，提供安全及数据完整性保障](https://zh.wikipedia.org/wiki/%E5%82%B3%E8%BC%B8%E5%B1%A4%E5%AE%89%E5%85%A8%E6%80%A7%E5%8D%94%E5%AE%9A)

    包含记录层（Record Layer）和传输层
        记录层协议:
            确定传输层数据的封装格式
        传输层安全协议:
            使用X.509认证
            之后利用非对称加密演算来对通信方做身份认证
            身份认证成功后交换对称密钥作为会谈密钥（Session key）
            这个会谈密钥是用来将通信两方交换的数据做加密，保证两个应用间通信的保密性和可靠性，使客户与服务器应用之间的通信不被攻击者窃听。

    主从式架构模型(Client/Server, C/S)

    优势
        是与高层的应用层协议（如HTTP、FTP、Telnet等）无耦合
        应用层协议能透明地运行在TLS协议之上，由TLS协议进行创建加密通道需要的协商和认证
        应用层协议传送的数据在通过TLS协议时都会被加密，从而保证通信的私密性

[生成证书](https://www.ibm.com/support/knowledgecenter/zh/SSWHYP_4.0.0/com.ibm.apimgmt.cmc.doc/task_apionprem_gernerate_self_signed_openSSL.html)
