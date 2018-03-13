# 实现

1 [核心技术](https://draveness.me/docker)

Namespaces

    Linux 提供的用于分离进程树、网络接口、挂载点以及进程间通信等资源的方法
    七种(创建新的进程时设置)：
        CLONE_NEWCGROUP
        CLONE_NEWIPC
        CLONE_NEWNET
        CLONE_NEWNS
        CLONE_NEWPID
        CLONE_NEWUSER
        CLONE_NEWUTS

    进程
        表示一个正在执行的程序，也是在现代分时系统中的一个任务单元，可使用 ps 命令查看
                                                        | -> cron (定时器)
                | -> /sbin/init (内核初始化和系统配置) ->
        idle ->                                         | -> tty
                | -> kthreadd (管理和调度其它的内核进程) -> watchdog
        在容器里查看进程
            sudo docker run -it -d ubuntu # 后台运行 ubuntu 镜像
            sudo docker exec -it CONTAINERID /bin/bash # 交互模式登录容器
            ps -ef # 查看容器里的进程
                UID        PID  PPID  C STIME TTY          TIME CMD
                root         1     0  0 06:07 pts/0    00:00:00 /bin/bash
                root         9     0  0 06:08 pts/1    00:00:00 /bin/bash
                root        18     9  0 06:08 pts/1    00:00:00 ps -ef
        容器成功将容器内的进程与宿主机器中的进程隔离
                                                                            | -> /bin/bash
            /sbin/init -> dockerd -> docker-containerd -> docker-containerd-shim
                                                                            | -> ps -ef
    网络
        容器必须能与外界通信才能发挥作用，因此 docker 为每个容器提供了单独的网络命名空间
        四种网络模式
            Host
            Container
            None
            Bridge
                默认模式
                会为容器设置IP地址，连接到虚拟网桥 Docker0
                每一个容器在创建时都会创建一对虚拟网卡，两个虚拟网卡组成了数据的通道，其中一个会放在创建的容器中，另一个会加入到名为 docker0 网桥中
                网桥 docker0 通过 iptables 中的配置与宿主机器上的网卡相连，所有符合条件的请求都会通过 iptables 转发到 docker0 并由网桥分发给对应的机器
                当有 Docker 的容器需要将服务暴露给宿主机器，就会为容器分配一个 IP 地址，同时向 iptables 中追加一条新的规则
                brctl show
        libnetwork
            Sandbox
                容器的网络栈配置，包括容器的接口、路由表和 DNS 设置
            Endpoint
                存在于 Sandbox 中，可以有一个或多个
            Network
    挂载点(目录)
        如果一个容器需要启动，那么它一定需要提供一个根文件系统（rootfs），容器需要使用这个文件系统来创建一个新的进程，所有二进制的执行都必须在这个根文件系统中
        为了保证当前的容器进程没有办法访问宿主机器上其他目录，我们在这里还需要通过 libcontainer 提供的 pivot_root 或者 chroot 函数改变进程能够访问个文件目录的根节点

CGroups(Control Groups)

    物理资源上的隔离，比如 CPU 、内存、磁盘 I/O 和网络带宽
    每一个 CGroup 都是一组被相同的标准和参数限制的进程，不同的 CGroup 之间是有层级关系的，也就是说它们之间可以从父类继承一些用于限制资源使用的标准和参数
    四种功能
        资源限制(resource limiting)
        优先级(prioritization)
        计数(accounting)
        控制(control)
    Linux 使用文件系统来实现 CGroup
    查看当前CGroup有哪些子系统：
        lssubsys -m
            cpuset /sys/fs/cgroup/cpuset
            cpu,cpuacct /sys/fs/cgroup/cpu,cpuacct
            blkio /sys/fs/cgroup/blkio
            memory /sys/fs/cgroup/memory
            devices /sys/fs/cgroup/devices
            freezer /sys/fs/cgroup/freezer
            net_cls,net_prio /sys/fs/cgroup/net_cls,net_prio
            perf_event /sys/fs/cgroup/perf_event
            hugetlb /sys/fs/cgroup/hugetlb
            pids /sys/fs/cgroup/pids
    新建CGroup：
        只需要在想要分配或者限制资源的子系统下面创建一个新的文件夹，然后这个文件夹下就会自动出现很多的内容
        如果你在 Linux 上安装了 Docker，你就会发现所有子系统的目录下都有一个名为 docker 的文件夹
            ls /sys/fs/cgroup/cpu
                ...
                docker
                ...
            ls /sys/fs/cgroup/cpu/docker/
                ...
                CONTAINERID # 容器ID作为目录名，启动这个容器时，Docker 会为这个容器创建一个与容器标识符相同的 CGroup
                ...
            cat /sys/fs/cgroup/cpu/docker/CONTAINERID/tasks
                4686 # 存储着属于当前控制组的所有进程的 pid，作为负责 cpu 的子系统
            cat /sys/fs/cgroup/cpu/docker/CONTAINERID/cpu.cfs_quota_us
                -1 # CPU 限制，如果是50000，表示不能超过50%；-1表示无限制
    Docker 关闭掉正在运行的容器时，Docker 的子控制组对应的文件夹也会被 Docker 进程移除，Docker 在使用 CGroup 时其实也只是做了一些创建文件夹改变文件内容的文件操作

UnionFS

    一种为 Linux 操作系统设计的用于把多个文件系统『联合』到同一个挂载点的文件系统服务

    每一个镜像层或者容器层都是 /var/lib/docker/ 目录下的一个子文件夹；在 Docker 中，所有镜像层和容器层的内容都存储在 /var/lib/docker/overlay2/diff/ 目录中(在ubuntu16.04普通用户有sudo也不能访问)

    构建镜像(镜像也是一个文件)

    存储驱动
        Docker 使用了一系列不同的存储驱动管理镜像内的文件系统并运行容器，这些存储驱动与 Docker 卷（volume）有些不同，存储引擎管理着能够在多个容器之间共享的存储
        Docker 中的每一个镜像都是由一系列只读的层组成的，Dockerfile 中的每一个命令都会在已有的只读层上创建一个新的层
        当镜像被 docker run 命令创建时就会在镜像的最上层添加一个可写的层，也就是容器层，所有对于运行时容器的修改其实都是对这个容器读写层的修改

    AUFS(升级版)
        提供更优秀的性能和效率
        能够将不同文件夹中的层联合（Union）到了同一个文件夹中，这些文件夹在 AUFS 中称作分支，整个『联合』的过程被称为联合挂载（Union Mount）

    其它存储驱动，包括 aufs、devicemapper、overlay2、zfs 和 vfs 等等
    在最新的 Docker 中，overlay2 取代了 aufs 成为了推荐的存储驱动，但是在没有 overlay2 驱动的机器上仍然会使用 aufs 作为 Docker 的默认驱动

    查看当前使用的存储驱动
        sudo docker info | grep Storage
            Storage Driver: overlay2
