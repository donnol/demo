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
