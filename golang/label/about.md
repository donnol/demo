# label 标签

## [Golang label](https://blog.csdn.net/qq_27682041/article/details/78765779)

分别是 break label 和 goto label 、continue label

这 3 个在多级嵌套语句中，非常有用

break label

    break的跳转标签(label)必须放在循环语句for前面，并且在break label跳出循环不再执行for循环里的代码。
    当我们把标签定义在break的下面时，我们会发现运行时报错。
    只能用于for循环，不能和switch使用

    结合 break 理解：
        break在for循环中执行就是跳出循环，执行循环后面的代码

goto label

    goto label的label(标签)既可以定义在for循环前面,也可以定义在for循环后面，当跳转到标签地方时，继续执行标签下面的代码

continue label

    跳出当前该次的循环圈，在Go编程语言中的continue语句有点像break语句。不是强制终止，只是继续循环下一个迭代发生，在两者之间跳过任何代码

    结合 continue 理解：
        continue在for循环中执行就是开始下一次循环，跳过当前循环剩余代码
