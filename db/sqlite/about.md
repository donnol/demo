# sqlite 数据库

1 介绍

    一个数据库就是一个文件

2 安装

    ubuntu 16.04 -> sudo apt install sqlite3

3 使用

    sqlite3 demo.db

4 事务

    锁：
        逐步上升机制：未加锁(UNLOCKED)、共享 (SHARED)、保留(RESERVED)、未决(PENDING)和排它(EXCLUSIVE)
        每个数据库连接在同一时刻只能处于其中一个状态
        读数据时，获得共享锁，允许其它连接继续获得共享锁，未全部释放前，都不能写
        写数据时，先获得保留锁（同一时刻只有一个），可与共享锁共存，在内存缓冲区修改数据
        要提交时，将保留锁提升为未决锁，其它连接不能再获得新的共享锁，已有共享锁继续工作，直到所有共享锁全部释放后，未决锁提升为排它锁，完成对数据库的修改
    死锁：
        当其中一个连接拥有未决锁，而其它连接在尝试将共享锁提升为保留锁时
    三种事务：
        DEFERRED 默认，开启事务时不会获取任何锁 ，当它读时，获取共享锁，当它写时，获取保留锁
        IMMEDIATE 开启事务时尝试获取保留锁，成功的话，其它连接不能写，只能读；会阻止其它连接的 BEGIN IMMEDIATE 或者 BEGIN EXCLUSIVE 命令，返回 SQLITE_BUSY 错误；当修改之后，必须等待其它读事务完成，才能提交
        EXCLUSIVE 开启事务时尝试获取排它锁，成功的话，其它连接不能再读写，直到该事务提交

5 WAL-Write Ahead Logging

    在引入 WAL  前，使用 rollback journey，在数据更新前将原记录复制到某个地方，然后直接修改数据库记录，如果成功提交，将复制删除，如果提交失败，将复制重新写回数据库
    在 3.7.0 后引入 WAL
    在事务修改数据时，将修改记录到wal文件里，然后提交事务，读数据时，先到wal查找，有则返回，无则到数据库查找；当wal文件积累了一定量后，将修改刷到数据库里

[官网](https://www.sqlite.org/)
[源码分析](http://huili.github.io/index.html)
[WAL](https://www.sqlite.org/wal.html)
