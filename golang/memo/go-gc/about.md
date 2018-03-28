# 垃圾收集

1 为什么 Go 不实现分代和紧凑 gc ？

    分代：
        假设分配在一个程序中的大部分值很快变得不会用到, 所以分代 GC 有一个优势就是可以花更多的时间查看最近分配的对象
        为什么不？
            Go 的许多对象是直接在程序栈(stack)上分配的。 Go 编译器使用逃逸分析(escape analysis)来查找那些在编译时生命周期就已知的对象, 将它们分配到堆栈而不是垃圾收集的内存中
    紧凑：
        避免碎片以及可以使用简单有效的凹凸分配器(bump allocator)
        为什么不？
            Go 运行时使用的是现代基于 tcmalloc 的方案，基本没有碎片问题

[来源](https://lingchao.xin/post/why-golang-garbage-collector-not-implement-generational-and-compact-gc.html)

2 tcmalloc

    现代化内存分配器
    基本特征：对抗内容碎片，在多核处理器能够扩展(scale)
    内存分配速度是 glibc2.3 中实现的 malloc 的数倍
    一个 Page 内：
        定长:
            用隐式 Freelist 进行对象分配 -- 最大化内存利用率，最小化分配时间
            Freelist:
                将 4KB 的内存划分为 16 字节的单元，每个单元的前8个字节作为节点指针(指针在分配出去之后就是数据区了，就不占空间了，只有待分配的时候才是指针，等于不需要额外空间)，指向下一个单元。初始化的时候把所有指针指向下一个单元；分配时，从链表头分配一个对象出去；释放时，插入到链表。
                由于链表指针直接分配在待分配内存中，因此不需要额外的内存开销，而且分配速度也是相当快。
        变长：
            简化成多种定长记录的分配
            如需要 7 字节，就分配 8 字节，向上取整
            不可避免地，这样会产生碎片，因此使用 8, 16, 32, 48, 64, 80 规则(除开 8,16 外，后面的以 16 字节增加)，使得碎片尽可能小
    大于一个 Page：
        多个连续 Page 组合成一个 Span(在 Span 里面会记录第一个 Page 的编号，以及 Page 的数量)
        大对象直接分配 Span ，小对象在 Span 中分配 Object
        不同数量的 Page ，组成不同大小的 Span ，如包含 1 个 Page 的 Span 有 n 个，2 个 Page 的 Span 有 m 个……
        初始时只有 128 Page 的 Span，如果要分配 1 个 Page 的 Span，就把这个 Span 分裂成两个，1 + 127，把127再记录下来
        Span 的合并：
            释放 Span 时，需要将前后的空闲 Span 进行合并，当然，前提是它们的 Page 要连续
            记录从 Page 到 Span 的映射：
                为了较小的空间开销，使用 RadixTree(压缩过的前缀数 trie) 数据结构
    总的分配流程：
        每个线程都一个线程局部的 ThreadCache，按照不同的规格，维护了对象的链表；如果ThreadCache 的对象不够了，就从 CentralCache 进行批量分配；如果 CentralCache 依然没有，就从PageHeap申请Span；如果 PageHeap没有合适的 Page，就只能从操作系统申请了。
        在释放内存的时候，ThreadCache依然遵循批量释放的策略，对象积累到一定程度就释放给 CentralCache；CentralCache发现一个 Span的内存完全释放了，就可以把这个 Span 归还给 PageHeap；PageHeap发现一批连续的Page都释放了，就可以归还给操作系统。

[来源](https://zhuanlan.zhihu.com/p/29216091)

[另外](http://kouucocu.lofter.com/post/1cdb8c4b_50f6300)

[Golang 实现 RadixTree](https://github.com/armon/go-radix/blob/master/radix.go)
