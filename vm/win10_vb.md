# 当你要在 win10 使用 virtualbox 安装 64 位系统的时候

1.检查电脑设置，确认 BIOS 中已经开启虚拟化支持

    Intel Visualization

2.win+Q 打开 Hyper-V 管理器，检查 hyper-v 管理服务状态，停止此服务

3.计算机管理——服务 中 禁用以 Hyper-V 开头的 8 个服务

4.以管理员权限运行如下命令 bcdedit /set hypervisorlaunchtype off

5.最后，一定要重新启动电脑，以上设置才能生效

## 当添加 ubuntu 虚拟机后，无法挂载共享目录

[from](https://askubuntu.com/questions/573596/unable-to-install-guest-additions-cd-image-on-virtual-box)

1 安装工具

    sudo apt-get install virtualbox-guest-utils

2 重启

3 挂载目录

    // go 在vbox设置，/mnt/go 在虚拟机系统目录里
    sudo mount.vboxsf go /mnt/go
    当设置共享文件夹为只读后，挂载时出现错误:mount.vboxsf: mounting failed with the error: Protocol error
    [解决办法1:将挂载点目录名改成和共享文件夹名字不一样 /mnt/go -> /mnt/share (试了后还是不行)](http://www.virtualbox.org/ticket/2265)
    解决办法2：在第一步后，修改vbox的共享文件夹 go -> vbshare ，再运行 sudo mount.vboxsf vbshare /mnt/share/ 或者 sudo mount -t vboxsf vbshare /mnt/share/

4 启动虚拟机系统时自动挂载

    在文件 /etc/fstab 后面添加
    vbshare /mnt/share    vboxsf ro,gid=1000,uid=1000,auto 0 0
    [在文件 /etc/modules 添加 vboxsf, 可选择性添加 vboxguest ](https://askubuntu.com/questions/365346/virtualbox-shared-folder-mount-from-fstab-fails-works-once-bootup-is-complete)
    在文件 /etc/rc.local 的 exit 0 前面添加
    mount.vboxsf vbshare /mnt/share/ 或者 mount -t vboxsf vbshare /mnt/share/

## SSH 连接虚拟机系统

需要在网络里添加一个网卡，设置 连接方式 为 NAT 模式

1 在 vbox 里添加 网络-高级-端口转发规则

    名称：ssh // 可以是任意名字，当有多个规则时作为区分
    协议: tcp
    主机IP: // 留空即可
    主机端口：3022 // 在ssh连接时使用 -p 指定
    子系统IP：// 也是留空即可
    子系统端口：22 // 虚拟机系统开启ssh服务时配置文件里的端口，ubuntu可在 /etc/ssh/ssh_config 文件里指定

2 在本地连接 ssh -p 3022 jd@127.0.0.1 // jd 替换为虚拟机系统的用户名

## 科学上网

1 本地已经有 ss ，虚拟机通过本地的 ss 代理

    ip route show 拿到 default via 后面的 虚拟路由器IP地址 vm-route-ip
    http_proxy=http://vm-route-ip:1080 ping www.google.com (这个试过了但是不行)

    需要在网络里添加一个网卡，设置 连接方式 为 桥接模式
    在宿主机 windows 上运行 shadowsocks.exe 并勾选“允许局域网连接”
    使用桥接方式运行虚拟机（这时虚拟机与宿主处于同一个局域网）
    进入ubuntu系统，System Settings – Network – Network proxy勾选Manual（手动）,地址全部填宿主机IP（局域网网段），设置好代理端口（可在windows下的shadowsocks查看，一般为默认1080）
    至此，使用浏览器访问www.google.com，就可以成功

    终端使用，需要设置环境变量 ALL_PROXY=socks5://[宿主机局域网ip地址]:1080 (这个方法好像不行哦，还是需要使用 polipo )
    或者
    也可以使用 [polipo](http://droidyue.com/blog/2016/04/04/set-shadowsocks-proxy-for-terminal/index.html) 将 socks 协议转换成 http 协议

    检验：
        curl ip.gs
        http_proxy=http://127.0.0.1:8123 curl ip.gs
        ALL_PROXY=socks5://192.168.1.120:1080 wget http://google.com
        如果连接失败，加 -i 可打印更多信息
    注意：
        测试连接的时候 http_proxy=http://127.0.0.1:8123 ping www.google.com 无反应，需要使用 curl 或者 wget 才可以

2 最后，写入 ~/.bashrc

    export http_proxy=http://127.0.0.1:8123
    export https_proxy=https://127.0.0.1:8123
    或者
    export ALL_PROXY=socks5://192.168.1.120:1080 (单纯设置这个的时候并不能成功，还是需要上面的设置)

## 备份与恢复

[关键是要保留 vdi 文件和快照的顺序(\*.vbox 里面 HardDisks 标签)](http://www.cnblogs.com/zjutlitao/p/5132610.html)

## 网络服务

在宿主机查看虚拟机

    /mnt/c/Oracle/VirtualBox/VBoxManage.exe guestproperty

get v-ubuntu "/VirtualBox/GuestInfo/Net/0/V4/IP"

在虚拟机里查看，使用虚拟机系统自带工具

    如在linux，可以 ifconfig, ip

1 宿主机访问虚拟机里运行的服务，方便使用宿主机的浏览器调试

    vbox 的 NAT 网卡里面添加多一条端口转发规则 8520 -> 8520
    建立规则后，在虚拟机运行过程中，本地不能再使用这个端口

2 虚拟机访问宿主机里运行的服务

    ？？估计是 polipo 配置的问题

        jd@jd-ubuntu:~/share/src/demo/server$ curl http://192.168.1.120:8510
        <!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">
        <html><head>
        <title>Proxy error: 502 Server dropped connection.</title>
        </head><body>
        <h1>502 Server dropped connection</h1>
        <p>The following error occurred while trying to access <strong>http://192.168.1.120:8510/</strong>:<br><br>
        <strong>502 Server dropped connection</strong></p>
        <hr>Generated Fri, 15 Dec 2017 16:40:07 CST by Polipo on <em>jd-ubuntu:8123</em>.
        </body></html>
    TODO 改天，新建一个虚拟机试下

3 两台虚拟机之间的通信

    v1 端口转发
        虚拟机 3306 -> 宿主机 3006
    v2 端口转发
        虚拟机 3306 -> 宿主机 4006
    v1 -> v2
        宿主机ip(xxx.xxx.xxx.xxx):4006
    v2 -> v1
        宿主机ip(xxx.xxx.xxx.xxx):3006

[官方文档-虚拟网络介绍](http://www.virtualbox.org/manual/ch06.html)

## 文件传输

1 scp 传输

    scp -P3022 /local/file username@remotehost:/path/to/save/

2 sftp 传输

    sftp username@remotehost:/path/to/go -P3022
    get xxx
    put xxx
    连接时出现了 Permission denied, please try again. 错误
        检查密码是否正确
    连接成功后，get file 时出现 Couldn't open local file "xxx" for writing: Permission denied
        检查一下是否拥有所在主机目录的读写权限

## 复制粘贴

[参考](https://askubuntu.com/questions/63420/how-to-fix-virtualboxs-copy-and-paste-to-host-machine)

1 安装增强功能(挂载共享目录时已安装)

2 关闭虚拟机，在虚拟机设置里

    设备-共享粘贴板-双向

    控制-设置-存储
        控制器SATA-勾选"使用主机输入输出(I/O)缓存"
        控制器SATA-点击***.vdi-勾选"固态驱动器"
    重启虚拟机

3 在虚拟机系统内安装

    sudo apt install xorg-video-abi-20 xserver-xorg-core
    sudo apt install virtualbox-guest-x11
    VBoxClient --clipboard
    出现错误: VBoxClient: Failed to connect to the VirtualBox kernel service
        没有加 sudo
    正确运行：sudo VBoxClient --clipboard

4 要使用 ctrl+c 和 ctrl+v 时

    使用 ctrl+shift+c 和 ctrl+shift+v 替代

## 桌面版 ubuntu 可选择安装 vscode

[参考](https://code.visualstudio.com/docs/setup/linux)

    curl https://packages.microsoft.com/keys/microsoft.asc | gpg --dearmor > microsoft.gpg
    sudo mv microsoft.gpg /etc/apt/trusted.gpg.d/microsoft.gpg
    sudo sh -c 'echo "deb [arch=amd64] https://packages.microsoft.com/repos/vscode stable main" > /etc/apt/sources.list.d/vscode.list'
    sudo apt-get update
    sudo apt-get install code # or code-insiders

## 启用 USB 设备后，宿主机鼠键无响应

## 虚拟机无法接收鼠键输入
