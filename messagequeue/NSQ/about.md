# NSQ

简介

    一个由 Go 编写的消息队列，使用 MIT 开源许可证
    分布式、去中心化的实时消息平台
    没有单点错误，容错，高可靠

[源码](https://github.com/nsqio/nsq)

    获取 go get -u github.com/nsqio/nsq/...
    依赖 dep ensure
    安装
        make
        sudo make install
    启动
        nsqd

nsqd 包

    64 位原子变量必须是结构体的第一个 field，因为这样在 32 位机器上也能正确地对齐

    管道类型 Channel

        一个 topic 下可以有多个管道，每个管道包含多个订阅者
        使用原子操作，切换暂停/开始 状态，记录消息数量，重新排队数量
        使用bufferPool(缓冲池)
        支持延迟发送
        重置发送中的消息的超时
        丢弃发送中的消息

    guidFactory 生产 guid

        生成更大，时间更靠前的 guid

    主题类型 Topic
