# 源码阅读笔记

## [bolt_unix.go](https://github.com/boltdb/bolt/blob/master/bolt_unix.go#L25:6)

### flock 和 funlock 函数

```go
// flock acquires an advisory lock on a file descriptor.
// flock 获取文件描述符的意向锁
func flock(db *DB, mode os.FileMode, exclusive bool, timeout time.Duration) error {
    ...
    flag := syscall.LOCK_SH // 共享锁
    if exclusive {
        flag = syscall.LOCK_EX // 排它锁
    }
    // 系统调用 syscall.Flock，syscall.LOCK_NB 表示非阻塞
    syscall.Flock(int(db.file.Fd()), flag|syscall.LOCK_NB)
    ...
}

// funlock releases an advisory lock on a file descriptor.
// funlock 释放文件描述符的意向锁
func funlock(db *DB) error {
    return syscall.Flock(int(db.file.Fd()), syscall.LOCK_UN) // 将标识设为 syscall.LOCK_UN
}
```

[flock 系统调用](https://blog.csdn.net/lqt641/article/details/54605920)

![进程与其打开的文件](https://img-blog.csdn.net/20170118225918700?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvbHF0NjQx/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
进程使用 flock 对已打开文件加文件锁时，是加在了上图中间蓝色部分的文件表项上

### mmap 和 munmap 函数

```go
// mmap memory maps a DB's data file.
// 映射一个数据库文件到内存
func mmap(db *DB, sz int) error {
    ...
    // 映射文件到内存
    syscall.Mmap(int(db.file.Fd()), 0, sz, syscall.PROT_READ, syscall.MAP_SHARED|db.MmapFlags)
    ...
}

// munmap unmaps a DB's data file from memory.
// 从内存中删掉数据库文件的映射
func munmap(db *DB) error {
    ...
    syscall.Munmap(db.dataref)
    ...
}
```

[mmap 系统调用](http://lib.csdn.net/article/linux/62126)

将磁盘文件映射到进程的虚拟地址空间, 通过对这段内存的读取和修改，来实现对文件的读取和修改,而不需要再调用 read，write 等操作
