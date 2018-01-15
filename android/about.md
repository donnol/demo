# 安卓系统

1 像普通服务器一样登录安卓手机终端

    通过手机应用 Android Play 安装 Termux
    打开 Termux 登录终端，安装 OpenSSH
        pkg install openssh
    启动
        sshd
    查看日志
        logcat -s 'syslog:*'
    退出日志
        ?
    连接
        ssh user@host -p 8022
    出现错误
        Permission denied (publickey,keyboard-interactive).
    在本地运行 ssh-keygen -t rsa 生成密钥
    将本地公钥 ~/.ssh/id_rsa.pub 复制到手机里的 .ssh/authorized_keys 文件
    手机端退出终端前，必须手动关闭 sshd(pkill sshd)，否则连接会一直存在

[官网](https://termux.com/)
[SSH](https://wiki.termux.com/wiki/SSH)
[中文教程](http://blackwolfsec.cc/2016/12/10/termux/)
[中文 SSH](https://www.findhao.net/easycoding/1652)
