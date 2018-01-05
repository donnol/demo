# 关于 redis

1 安装

    sudo apt install redis-server redis-tools

2 登录

    redis-cli

3 设置密码

    sudo vim /etc/redis/redis.conf
    将 requirepassword 行的注释去掉，并在后面设置密码
    重新启动redis-server
    登录后，使用 auth pass 校验

4 读写

    字符串类型:
        set a 1
        get a
    哈希表类型:
        hset c 1 '1'
        hget c 1
    列表类型:
        lpush la 1
        lpop la
    集合类型:
        sadd sa 1
        spop sa
    有序集合类型:
        zadd za 1 1
        zscan za 1
    pub/sub 类型:
        client1 > subscribe msg
        client2 > publish msg 'msg'
    事务类型:

5 切换数据库

    SELECT [0~15]

[内部数据结构分析](http://zhangtielei.com/posts/blog-redis-dict.html)
