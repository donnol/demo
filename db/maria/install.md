# ubuntu 16.04 安装 mariadb

1 安装

    sudo apt install mariadb-client mariadb-server

    默认安装的是10.0，如果想要最新的10.2，请如下操作
        sudo apt-get install software-properties-common
        sudo apt-key adv --recv-keys --keyserver hkp://keyserver.ubuntu.com:80 0xF1656F24C74CD1D8
        在 /etc/apt/sources.list.d/ 目录下创建 MariaDB.list ，然后添加一下内容到文件
            # MariaDB 10.2 repository list - created 2017-12-27 07:16 UTC
            # http://downloads.mariadb.org/mariadb/repositories/
            deb [arch=amd64,i386] http://mirrors.tuna.tsinghua.edu.cn/mariadb/repo/10.2/ubuntu xenial main
            deb-src http://mirrors.tuna.tsinghua.edu.cn/mariadb/repo/10.2/ubuntu xenial main
        最后，运行
            sudo apt update
            sudo apt install mariadb-server
        使用时，出错：ERROR 1524 (HY000): Plugin 'unix_socket' is not loaded
        unix_socket 表明 mariadb 将使用本机用户信息作为登录校验，而不是传统的 mysql_native_password 密码校验
        解决的话，需要使用 mysqld_safe --skip-grant-tables 登录 root 用户，安装插件
                INSTALL PLUGIN unix_socket SONAME 'auth_socket';

[安装指定版参考](https://downloads.mariadb.org/mariadb/repositories/#mirror=tuna&distro=Ubuntu&distro_release=xenial--ubuntu_xenial&version=10.2)
[unix_socket 参考](https://mariadb.com/kb/en/library/authentication-plugin-unix-socket/)
[unix_socket 错误解决参考](https://askubuntu.com/questions/705458/ubuntu-15-10-mysql-error-1524-unix-socket)

2 安全配置

    mysql_secure_installation
        set root passwd?
        remove anonymous users?
        Disallow root login remotely?
        Remove test database and access to it?
        Reload privilege tables now? 重新载入权限表? 选 Y，上面的配置才会 生效
