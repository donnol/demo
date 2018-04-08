# install

1 ubuntu 16.04 安装 postgresql-10

    添加 apt 源
    vim /etc/apt/sources.list.d/pgdg.list
    添加 deb http://apt.postgresql.org/pub/repos/apt/ xenial-pgdg main
    签名
        wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | \
        sudo apt-key add -
    sudo apt update
    sudo apt install postgresql-10
    登陆
    sudo -u postgres psql
    创建新库
    create database demo;
    创建角色
    create role jd LOGIN PASSWORD 'jd' REPLICATION CREATEDB;
    查看角色
    \du 或者 select * from pg_roles;
    新角色登陆
    psql -Ujd

2 安装 rdkit

    查看可用插件
    select * from pg_available_extensions;
    查看已有插件
    \dx 或 select * from pg_extension;
    安装 rdkit
    sudo apt-get install python-rdkit librdkit1 rdkit-data
    创建插件
    create extension rdkit;

[参考](http://www.rdkit.org/docs/Install.html)

3 docker 启动

    启动
        sudo docker run -it -p 5432:5432 postgres # 这里会启动容器
    获取容器名
        sudo docker ps # 最后一列 NAMES 的值，下面登录时使用 lucid_kilby 是我拿到的容器名
    psql 登录数据库
        sudo docker exec -it lucid_kilby psql -U postgres -d demo # 使用 exec，-d 指定数据库
    登录 bash
        sudo docker exec -it lucid_kilby bash
    配置文件目录
        /var/lib/postgresql/data/postgresql.conf
    如果想将本地配置挂载到容器里，可在启动容器时用 -v myconf:containerconf 指定
        sudo docker run -it -p 5432:5432 -v myconf:containerconf postgres # containerconf 一般是 /var/lib/postgresql/data/postgresql.conf
