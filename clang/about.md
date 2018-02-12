# C 语言

关键字(32 个)

    类型：
        char
        double
        enum: 枚举，值默认从零开始，也可以自定义，必须是整型常量
        float
        int
        long
        short
        signed
        struct
        union: 可以用来提高内存使用率，一个 union 只配置一个足够大的空间以来容纳最大长度的数据成员
            它维护足够的空间来存放多个数据成员中的“一种”，而不是为每一个数据成员配置空间，在 union 中所有的数据成员共用一个空间，同一时间只能储存其中一个数据成员
            所有的数据成员具有相同的起始地址
        unsigned
        void
    控制语句：
        for
        do
        while
        if
        else
        switch
        case
        default
        goto
        continue
        break
        return
    存储类型：
        auto: 声明自动变量，声明变量时默认是 auto
        extern: 声明变量是在其它文件声明的，全局变量和函数默认是 extern
        register: 声明寄存器变量，保存在CPU寄存器里，读取非常快，但是数量有限；超出限制后，转为auto
        static: 声明静态变量
            1 全局变量/函数的作用范围变为本源文件
            2 内存存储区域不再是栈，而是在静态存储区，跟全局变量一样
                系统在编译阶段就为这个变量所分配到的静态存储区的赋值了，所以
                    程序中的以下赋值语句
                        static int count = 10;
                    是不会运行的
            3 局部变量的默认初始化值为0，全局变量也具备这一属性
    其它：
        const: 声明只读变量，必须在定义时进行初始化，初始化后不能改变，跟#define一样，但是它是类型安全的
        sizeof: 计算数据类型长度
            与 strlen 的区别：sizeof是运算符，而strlen是函数；sizeof计算的是数据类型的大小，而strlen计算的是字符串的长度；sizeof的参数既可以是数据类型，也可以是变量，而strlen的参数只能是char*,而且必须是空字符结尾；sizeof返回值类型为unsigned，而strlen返回值为signed，因为它需要返回负数来表示出错情况
        typedef: 给数据类型取别名，起替代作用，用一些短小精悍的词替代一些复杂的类型
        volatile: 声明变量在程序执行中可被隐含地改变
            1 不允许编译器对与它有关的运算做任何优化
            2 变量可能会在程序外被改变，所以每次都必须从内存中读取，而不能把他放在cache或寄存器中重复使用
            用于：并行设备的硬件寄存器；中断服务子程序中会访问到的非自动变量；多线程应用中被几个任务共享的变量
    C99新添：
        restrict: 只能修饰一个指针，只要这个指针活着，这个指针独享这片内存，没有‘别人’可以修改这个指针指向的这片内存，所有修改都得通过这个指针来