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

1 在 vbox 里添加 网络-高级-端口转发规则

    名称：ssh // 可以是任意名字，当有多个规则时作为区分
    协议: tcp
    主机IP: // 留空即可
    主机端口：3022 // 在ssh连接时使用 -p 指定
    子系统IP：// 也是留空即可
    子系统端口：22 // 虚拟机系统开启ssh服务时配置文件里的端口，ubuntu可在 /etc/ssh/ssh_config 文件里指定

2 在本地连接 ssh -p 3022 jd@127.0.0.1 // jd 替换为虚拟机系统的用户名
