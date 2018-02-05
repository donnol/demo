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
