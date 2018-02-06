# 操作系统做些什么

1 进程管理

    当一个程序被加载进内存，就成为了一个进程
    可以分为四部分：
        栈: 包含了临时数据，如函数/方法参数、返回地址、局部变量
        堆: 运行时动态分配的内存
        数据: 全局和静态变量
        文本: 包括了程序计数器代表的当前活动、处理器寄存器的内容
    进程执行时，它会经历不同状态:
        开始：进程刚开启时的初始化状态
        准备：进程等待着可以被操作系统分配处理器，刚开始/中断后会进入到该状态
        等待：进程需要等待某个资源时，比如等待用户输入、文件可操作
        运行：操作系统调度器分配了处理器给进程时
        终止：进程完成了操作/被操作系统终止时，在这里等待被移除出内存
    进程控制块(PCB)：
        由操作系统维护的一个数据结构:
            PID: 进程 ID，整型，进程的标识，唯一
            State: 进程的当前状态
            Privileges：系统资源的基本权限
            Pointer：指向父进程的指针
            Program Counter：程序计数器，一个指向接下来将要执行的指令的指针
            CPU Registers：处理器寄存器，储存进程状态的各个寄存器
            CPU Scheduling Information：处理器调度信息，包括进程优先级等
            Memory Management Information：内存管理信息，包括页表，内存范围，段表等
            Accounting Information：统计信息，包括进程使用的处理器数量，时间范围，执行 ID 等
            IO Status Information：IO 状态信息，进程有关的 IO 设备列表
            etc...
    有关函数：
    有关工具：

2 线程和并发

    线程是进程里的一个执行流程，它有自己的程序计数器，知道自己接下来将要运行什么指令；有自己的系统寄存器，保存当前工作中的变量；还有一个包含执行历史的栈
    线程间可以共享一些信息，如代码段、数据段和打开的文件。当有一个线程改变了这些信息，其它线程都能看到这些改变
    优势：
        减少上下文切换开销
        允许多核
    分为：
        用户级线程
            优势：
                不需要内核权限
                跨平台
                专门调度
                快速创建和管理
            劣势：
                在大部分操作系统里，大部分系统调用是阻塞的
                不能获得和多进程一样的好处
        内核级线程
            优势：
                内核可以直接调度同一个进程里面的多个线程
                如果进程里的一个线程阻塞了，内核可以调度同一个进程里的其它线程
            劣势：
                创建和管理比用户级慢
                在一个进程里切换线程需要进入内核态
    有关函数：
    有关工具：

3 进程调度

    操作系统在进程调度队列里维护着所有的PCB
    对处于不同的状态的进程，操作系统会用不同的队列来维护它们。当它们的状态改变时，操作系统将它们移动到新的状态对应的队列
    有以下队列：
        Job queue：保存系统的所有进程
        Ready queue：保存在内存中，准备和等待被执行的进程，新进程总会被放入这个队列
        Device queues：由于 IO 设备不可操作而阻塞的进程
    操作系统可以使用不同的策略管理这些队列，如FIFO、Round Robin, Priority等
        FIFO: 先进先出
        Round Robin: ？
        Priority：优先级
    有关函数：
    有关工具：

4 内存管理

    将进程从内存和磁盘上来回移动
    有关函数：
    有关工具：
        free
        top

5 进程间通信

    共享内存
    消息解析
        header + body
    有关函数：
    有关工具：

6 IO 管理

    类别：
        Block Devices: 如硬盘，USB等
        Character Devices：串行端口，并行端口，声卡等
    通信方法：
        特别指令
        内存映射：内存和 IO 设备共享一个地址空间，如此，不用通过 CPU 即可传输数据，使用高速 IO 设备
        直接内存访问：允许 IO 模块直接读写内存；像键盘这样的慢设备会在传输每个字符之后产生一个中断，如果是硬盘这样的快设备，每个字符都要产生一个中断，就会消耗大部分 CPU 时间，这时，使用 DMA 就可以减少这部分消耗
            需要使用特殊硬件 DMAC，用于管理数据传输和系统总线访问
    有关函数：
    有关工具：

7 可视化

    数据可视
    桌面可视
    服务可视
    操作系统可视
    网络功能可视
    有关函数：
    有关工具：

8 分布式文件系统

    有关函数：
    有关工具：

9 分布式共享内存

    有关函数：
    有关工具：

10 云计算

    有关函数：
    有关工具：

[来源](https://codeburst.io/how-operating-systems-work-10-concepts-you-should-know-as-a-developer-8d63bb38331f)