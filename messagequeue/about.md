# 消息队列

主要特点是异步处理，主要目的是减少请求响应时间和解耦。

所以主要的使用场景就是将比较耗时而且不需要即时（同步）返回结果的操作作为消息放入消息队列。

同时由于使用了消息队列，只要保证消息格式不变，消息的发送方和接收方并不需要彼此联系，也不需要受对方的影响，即解耦和

[美团的文章](https://tech.meituan.com/mq-design.html)
