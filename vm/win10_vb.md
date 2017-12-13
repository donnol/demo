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

    // go 在vbox设置，/mnt/go 在系统目录里
    sudo mount.vboxsf go /mnt/go
