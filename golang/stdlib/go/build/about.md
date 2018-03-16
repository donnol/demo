# build

收集关于 Go 包的信息

    go path
        包含源码的目录树
    构建约束
        // +build
    二进制包
        //go:binary-only-package // 必须在文件顶部；必须紧贴着，没有空格
        // 必须有空行
        package mypkg
