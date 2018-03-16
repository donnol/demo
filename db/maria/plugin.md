# 插件

1 审计插件安装

    查看是否已安装
        show global variables like '%audit%';
    查看存放插件的路径
        show variables like 'plugin_dir';
    查看数据库版本
        select version();
    查看插件名字
        ls /plugin/path # 这里是审计插件
    安装插件
        install plugin server_audit soname 'server_audit.so';
    再次查看确认是否已安装
    启用插件
        set global server_audit_logging=on;
    设置审计日志路径
        set global server_audit_file_path='/path/to/log/';
    设置日志操作指令内容
        set global server_audit_events='QUERY_DDL,QUERY_DML';
    设置日志文件大小
        set global server_audit_file_rotate_size='200000000';
    设置日志数量
        set global server_audit_file_rotations='200';
    设置需要审计的用户
        set global server_audit_incl_users='root';
    设置免审计的用户
        set global server_audit_excl_users='jd';
    设置标识符
        set global server_audit_syslog_ident='mysql-server_auditing';

    为使设置不随重启而重置，请将以上配置写入配置文件 /etc/my.cnf 的 [mysqld] 下
