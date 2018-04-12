# cache

一般翻译为缓存，但代表的意思更接近'快取'

作为一个数据模型对象，它的特征

    命中率
        = 返回正确结果数/请求缓存次数
    最大元素（或最大空间）
        缓存中可以存放的最大元素的数量，一旦缓存中元素数量超过这个值（或者缓存数据所占空间超过其最大支持空间），那么将会触发缓存启动清空策略
    清空策略
        FIFO(先进先出策略，最先进入缓存的数据在缓存空间不够的情况下（超出最大元素限制）会被优先被清除掉)
        LFU(最少使用策略，无论是否过期，根据元素的被使用次数判断，清除使用次数较少的元素释放空间)
        LRU(最近最少使用策略，无论是否过期，根据元素最后一次被使用的时间戳，清除最远使用时间戳的元素释放空间)

缓存介质

    内存：将缓存存储于内存中是最快的选择，无需额外的I/O开销，但是内存的缺点是没有持久化落地物理磁盘
    硬盘：一般来说，很多缓存框架会结合使用内存和硬盘
    数据库：前面有提到，增加缓存的策略的目的之一就是为了减少数据库的I/O压力，如 Redis

缓存分类

    memcached缓存 分布式缓存
    redis缓存 分布式集群

[参考](https://tech.meituan.com/cache_about.html)

[通过 http 头字段控制缓存](http://www.alloyteam.com/2016/03/discussion-on-web-caching/)
[MDN](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Caching_FAQ)

[缓存更新](https://coolshell.cn/articles/17416.html)
[缓存一致性](http://www.infoq.com/cn/articles/cache-coherency-primer/)

## buffer

一般翻译为缓冲

[两者区别](https://www.zhihu.com/question/26190832)
