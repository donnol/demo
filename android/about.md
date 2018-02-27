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
        (音量+)+q 调出辅助输入栏，就会有 ESC TAB 等
    连接
        ssh user@host -p 8022
    出现错误
        Permission denied (publickey,keyboard-interactive).
    在本地运行 ssh-keygen -t rsa 生成密钥
    将本地公钥 ~/.ssh/id_rsa.pub 复制到手机里的 .ssh/authorized_keys 文件
    再次连接即可
    手机端退出终端前，必须手动关闭 sshd(pkill sshd)，否则连接会一直存在

[官网](https://termux.com/)
[SSH](https://wiki.termux.com/wiki/SSH)
[中文教程](http://blackwolfsec.cc/2016/12/10/termux/)
[中文 SSH](https://www.findhao.net/easycoding/1652)

2 使用 Go

[参考](http://rafalgolarz.com/blog/2017/01/15/running_golang_on_android/)

    安装
        pkg install golang
    测试
        go version
        go env
        添加文件 hello.go
            package main
            import "fmt"
            func main() {
                    fmt.Println("hello android.")
            }
        直接运行 go run hello.go
        或者编译后运行
            go build -o hello
            ./hello
    在主机编译好二进制文件，再传输到安卓机器使用
        GOOS=linux GOARCH=arm go build -o hello
        scp -P 8022 /local/hello user@host:/data/data/com.termux/files/home/hello
