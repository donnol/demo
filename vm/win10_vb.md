# 当你要在 win10 使用 virtualbox 安装 64 位系统的时候

1.检查电脑设置，确认 BIOS 中已经开启虚拟化支持

    Intel Visualization

2.win+Q 打开 Hyper-V 管理器，检查 hyper-v 管理服务状态，停止此服务

3.计算机管理——服务 中 禁用以 Hyper-V 开头的 8 个服务

4.以管理员权限运行如下命令 bcdedit /set hypervisorlaunchtype off

5.最后，一定要重新启动电脑，以上设置才能生效
