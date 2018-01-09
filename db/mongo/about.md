# mongo

1 [安装](https://docs.mongodb.com/manual/tutorial/install-mongodb-on-ubuntu/)

    旧版本
    sudo apt install mongodb-server mongodb-clients mongodb
    新版本
    sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 2930ADAE8CAF5059EE73BB4B58712A2291FA4AD5
    echo "deb [ arch=amd64,arm64 ] https://repo.mongodb.org/apt/ubuntu xenial/mongodb-org/3.6 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-3.6.list
    sudo apt udpate
    sudo apt-get install -y mongodb-org=3.6.1 mongodb-org-server=3.6.1 mongodb-org-shell=3.6.1 mongodb-org-mongos=3.6.1 mongodb-org-tools=3.6.1

2 登录

    启动
        sudo service mongod start
    查看
        ps aux | grep mongo
    安装旧版本后 mongo 出现错误 -- 因为 mongod 启动失败
        MongoDB shell version: 2.6.10
        connecting to: test
        2018-01-08T14:49:31.074+0800 warning: Failed to connect to 127.0.0.1:27017, reason: errno:111 Connection refused
        2018-01-08T14:49:31.074+0800 Error: couldn't connect to server 127.0.0.1:27017 (127.0.0.1),
        connection attempt failed at src/mongo/shell/mongo.js:146
        exception: connect failed
    安装新版本，直接用 mongo 进入 shell

3 使用

    use test // 使用数据库 test
    db.collection.insert({'name': 'jd'}) // 往 test 数据库的集合 collection 里添加文档
    db.collection.find() // 在集合里查找文档
    db.collection.find({'name': 'jd'}) // 查找指定文档
    db.collection.find({'name': {$gt: 'ja'}})  // 范围查找
    db.colleciton.find({'name': 'jd', 'age': 18}) // AND 查找
    db.collection.find({$or: [{'name': 'jd'}, {'name': 'ja'}]}) // OR 查找
    db.collection.find().sort({'name': 1}) // 排序
    db.collection.update({'name': 'jd'}, {$set: {'name': 'jj'}}) // 更新
    db.collection.update({'name': 'jd'}, {$set: {'name': 'jj'}}, {multi: true}) // 更新多行
    db.collection.remove({'name': 'jj'}) // 删除
    db.collection.remove({'name': 'jj'}, {justone: true}) // 删除一行
    db.collection.remove({}) // 删除全部
    db.collection.drop() // 删除集合
    db.collection.aggregate([{$match: {'name': 'jd'}}, {$group: {'_id': '1', 'count': {$sum: 1}}}]) // 聚集函数
    db.collection.createIndex({'name': 1}) // 创建索引
    db.collection.createIndex({'name': 1, 'age': 1}) // 组合索引
    db.collection.createIndex({'name': 1}, {unique: true}) // 唯一索引
    db.collection.dropIndex({'name': 1}) // 删除索引
    db.collection.getIndexes() // 查看集合索引
    db.collection.find().explain() // 查询计划

4 数据类型

    文档
        BSON Binary JSON 的简称，是一个JSON文档对象的二进制编码格式
        扩展了 JSON，如支持日期类型、BinDate类型等
        三个特点：轻量性、可遍历性、高效性
        文档中的键/值对是有序的
        区分类型和大小写
    ObjectId(_id)
        相当于主键，全局唯一
        4 字节时间戳 + 3 字节机器唯一标识码 + 2 字节进程 ID + 3 字节随机数
            > db.collection.find()
            { "_id" : ObjectId("5a540ca1ee279db6d2396eb7"), "name" : "jd" }
            5a540ca1 -> 4 字节时间戳
            ee279d -> 3 字节机器唯一标识码
            b6d2 -> 2 字节进程 ID
            396eb7 -> 3 字节随机数

5 [索引](http://lib.csdn.net/article/mongodb/53952)

    使用 B-tree (多路自平衡的搜索树)
        B+tree 只有叶节点存放数据(会有指针相连)，其余节点用来索引，而 B-tree 是每个索引节点都会有 Data  域。
        所以 B+tree 更适合用来存储外部数据，也就是所谓的磁盘数据(如 mysql )

[参考](https://segmentfault.com/a/1190000004690721)

6 MVCC

    3.0 后，使用 WiredTiger 作为默认引擎，通过该引擎的 MVCC 开始支持文档级别的锁

[参考](https://dawning7670.github.io/2017/10/17/MongoDB%E5%AD%98%E5%82%A8%E5%BC%95%E6%93%8E/)
[参考](https://draveness.me/mongodb-wiredtiger)

7 存储引擎

[参考](http://wudaijun.com/2016/06/mongodb-storage-engine/)

8 最新

    v3.6 主要提升安全性
