# 主从配置

环境：

    vbox 里跑两台 ubuntu 虚拟机，分别为 v1，v2，其中 v1 为 Master，v2 为 Slave
    mariadb 10.3

步骤：

    删除系统上已有的数据库(这步可以跳过)
        sudo apt remove mysql-*
        dpkg -l |grep ^rc|awk '{print $2}' |sudo xargs dpkg -P
        sudo rm -rf /etc/mysql/
    然后重新安装，参考 install.md

    v1 上：
        修改配置文件 /etc/mysql/my.cnf，添加以下内容
            # Master Settings
            log-bin
            server_id=1
            replicate-do-db=important # 指定一个数据库，这里是 important
        并把 bind-address 那行注释掉
        修改配置后记得重启
        登录数据库 mysql -uroot -p
        建立一个可以访问该数据库的用户(给从服务器使用)，这里是 slaveuser，密码是 slavepwd
            GRANT REPLICATION SLAVE ON *.* TO  'slaveuser'@'%' IDENTIFIED BY 'slavepwd';
            FLUSH PRIVILEGES;
            GRANT ALL PRIVILEGES ON important.* TO 'slaveuser'@'%' WITH GRANT OPTION;
            FLUSH PRIVILEGES;
        获取关键属性值 SHOW MASTER STATUS # 拿到 MASTER_LOG_FILE 和 MASTER_LOG_POS

    v2 上：
        修改配置文件 /etc/mysql/my.cnf，添加以下内容
            # Slave Settings
            server_id = 2
            replicate-do-db=important # 同 v1 配
        修改配置后记得重启
        给予用户权限，这里的 slaveuser 可以用其它名字吗？ TODO
            GRANT ALL PRIVILEGES ON important.* TO  'slaveuser'@'%' WITH GRANT OPTION;
            FLUSH PRIVILEGES;
        启动：
            STOP SLAVE;
            CHANGE MASTER TO MASTER_HOST= 'xxx.xxx.x.xxx', MASTER_PORT = xxx,MASTER_USER = 'slaveuser', MASTER_PASSWORD = 'slavepwd', MASTER_LOG_FILE = 'xxx-bin.000001', MASTER_LOG_POS = xxx;
            START SLAVE;
            SHOW SLAVE STATUS\G;
            第一行显示 Slave_IO_State: Waiting for master to send event ，那就是成功了

测试：

    v1 上：
        创建数据库 create database important
        在 important 库里新建表 t1 ，并添加数据
            create table t1 (id int not null, name varchar(256) not null);
            insert into t1 (id, name) values( 1, 'jd');
    v2 上：
        当 v1 上有数据变化时，v2 上就会出现相应的数据
            select * from t1;

    如果在 v2 上添加数据，如
        insert into t1 (id, name)values(2, 'jj');
    v1 上是不会有这条数据的，也就是说只有 v2 同步 v1 的数据，v1 不会同步 v2 的数据的

[详见](https://mariadb.com/kb/en/library/setting-up-replication/)
