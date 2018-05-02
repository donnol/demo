# 语言入门之一 关键字

熟悉关键字的使用和实现

关键字不能用于自定义名字，只能在特定语法结构中使用。

## 25 个

    break: 跳出
    default: 默认分支
    func: 声明函数
    interface: 接口
    select: 执行多个条件，选择满足条件的分支执行；如果同时有多个满足，则随机选取一个执行
    case: 分支
    defer: 延迟执行，在函数返回前执行
    go: 启动一个 goroutine
    map: 哈希表
    struct: 结构体
    chan: 管道
    else: 分支
    goto: 跳转
    package: 声明包
    switch: 选择条件满足的分支
    const: 常量
    fallthrough: 分支穿透
    if: 条件
    range: 遍历
        会复制一份数据保存到临时变量，注意slice和map的复制
        参考: https://garbagecollected.org/2017/02/22/go-range-loop-internals/
    type: 自定义类型
    continue: 继续下一次循环
    for: 开始一个循环
    import: 导入包
    return: 返回
    var: 变量

除此之外，如 int 和 true 等，属于预定义名字，主要对应內建的常量、类型和函数
