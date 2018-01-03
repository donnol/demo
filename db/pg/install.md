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
