# tool

iftop

    监听您指定的网络接口，并以 top 的样式呈现

    可以用于找出网络拥塞，测速和维持网络流量总量

    例如：sudo iftop

dig

    找到某域名所在服务器 ip 地址

    例如：dig A www.jdlau.com

ipcalc

    获取 ip 的子网掩码

    例如：ipcalc -b [ip]

nethogs

    当你想要快速了解是谁在吸取你的带宽的时候，nethogs 是个快速而简单的方法

    例如：sudo nethogs

vnstat

    它以守护进程在后台运行，因此可以实时地记录你的网络数据

    例如：vnstat

find

    查找文件

    例如：find test

objdump

    查看二进制文件信息

    例如：objdump [option] /path/to/binary

ftrace

    是一个 Linux 内核特性，它可以让你去跟踪 Linux 内核的函数调用

    # ftrace 比较难用，使用 trace-cmd 替代

    跟踪函数
        # __do_page_fault 是一个函数名字，当发生页面故障时调用
        sudo trace-cmd list -f __do_page_fault // 确认函数是否存在
        sudo trace-cmd record -p function -l __do_page_fault // 生成一个 trace.dat 文件
        sudo trace-cmd report // 显示 trace.dat 的内容
            trace-cmd-14940 [000] 883306.737031: function:             __do_page_fault
            trace-cmd: 进程名
            14940: 进程ID
            [000]: CPUID

     跟踪进程
        ps aux | grep [process] // 获得进程ID
        sudo trace-cmd record -p function -P 25314 // -P 指定进程ID

    出现提示
        Note: "entries" are the entries left in the kernel ring buffer and are not
        recorded in the trace data. They should all be zero.

strace

    跟踪进程执行时的系统调用和所接收的信号

    例如：strace nginx

unhide

    一个小巧的网络取证工具，能够发现那些借助 rootkit、LKM 及其它技术隐藏的进程和 TCP/UDP 端口

    例如：sudo unhide [options] proc

nmap

    一个网络探测工具和安全/端口扫描器
    发现网络上有哪些主机
    查看整个网络的信息
    提供关于目标机的进一步信息，包括反向域名，操作系统猜测，设备类型，和MAC地址

    例如：nmap 192.168.1.0/24

ifconfig

    查看网络信息

    例如：ifconfig

ip

    查看网络信息

    例如：ip

fio

    测试磁盘读写速度(iops)

    例如：sudo fio --bs=4k --ioengine=libaio --iodepth=32 --direct=1 --rw=randread --time_based --runtime=60 --refill_buffers --norandommap --randrepeat=0 --group_reporting --name=fio-read --size=100G --filename=/dev/vda

lftp

    文件传输程序，可以并行下载

    例如：lftp -e 'pget -n 5 -c http://kernel.org/pub/linux/kernel/v2.6/linux-2.6.22.2.tar.bz2'

iptables

    配置 Linux 内核 防火墙 的命令行工具(将会被 nftables 替代)
    iptables 可以检测、修改、转发、重定向和丢弃 IPv4 数据包
    iptables规则是数组式布局
    机械式的，一次处理一条规则

    例如：sudo iptables -I INPUT -s xxx.xxx.xxx.x/28 -j DROP # 参数化

nftables

    Netfilter
    nftables规则是链表式布局
    不是机械式的，是可编程的

    例如：nft add rule ip filter output tcp dport 80 drop # 语法

lscpu

    查看 CPU 信息

    例如：lscpu
    类似：cat /proc/cpuinfo # 信息更多

free

    查看内存信息

    例如：free -m
    类似：cat /proc/meminfo
    类似：dmidecode -t memory

lsblk

    查看磁盘和分区

    例如：lsblk
    类似：fdisk -l

lspci

    查看主板所有硬件槽信息

    例如：lspci

du

    统计文件占用的磁盘大小

    例如：du /path -h -t 1024000 --exclude=xxx # -h: 以 K 为单位展示; -t: 指定阈值; --exclude: 排除 xxx 文件/目录

ldd

    访问共享对象依赖关系(列出可执行文件依赖的动态文件，如 .so)

    例如：ldd a.out

showterm

    录制终端会话

    例如：showterm # 开始，结束时输入 exit 即可结束录制，并把视频上传到 http://showterm.io 网站上，暂时不支持video格式

[演示](http://showterm.io/7aeed561c67c7c852b09a)

ffmpeg

    音视频格式转换

    例如：ffmpeg -t 5 -ss 00:00:10 -i funny.mp4 out%04d.gif # 解压输入视频的视频帧，从第10秒开始，每5秒一帧

imagemagick

    将图片合并成gif动图

    例如：convert -delay 1x20 -loop 0 out*.gif animation.gif # 生成一副每秒20帧和循环无数次的动态GIF图片，其中 out*.gif 为输入图片，animation.gif 为输出动图

dnsmasq

    开源的轻量级DNS转发和DHCP、TFTP服务器，使用C语言编写

    例如：sudo dnsmasq --no-daemon

dig

    DNS 查询工具(Domain Information Groper)

    例如：dig jdscript.com

hardinfo

    查看机器硬件信息

    例如：hardinfo

Tripwire

    检查系统文件的完整性

    例如：sudo tripwire --init && sudo tripwire --check
