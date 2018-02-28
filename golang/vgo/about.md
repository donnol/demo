# vgo-Go 的版本控制

疑问

    1 在 go.mod 文件里 require 的包在 build 的时候不会下载到本地的吗？
        模块始终下载通过 HTTP 提供的 zip 归档文件
        会在 $GOPATH/src 目录里新建 v 目录，保存下载的文件
            v 目录包含 cache 目录和依赖包目录
            其中，cache 目录里保存各依赖包不同版本的信息
                info 文件记录包版本等信息
                mod 文件记录包依赖信息
                ziphash 文件记录 zip 归档文件的 hash 信息
                zip 文件是该版本的 zip 包
            依赖包目录里面是 zip 解压后的源码文件
    2 存储到 GOCACHE 环境变量设置的目录里吗？
        使用 vgo build -x
        可以看到 packagefile 指向的文件都在 $GOCACHE 目录里
        将 GOCACHE 的值临时设为 mytmp 目录后，可以看到更详细的构建过程

        整个编译流程，先将源文件(.go) compile 生成文件(.a)，然后再用 link 将生成的 .a 文件生成可执行文件
        构建时生成的临时文件会缓存到该目录，方便重用

导入兼容性规则

    如果旧包和新包具有相同的导入路径, 新软件包必须向后兼容旧软件包
    若要打破兼容性，请使用新的导入路径(在路径中加入语义版本)

    当你改变一个函数的行为时, 你也改变它的名字 -- Rob Pike

最小版本选择

    低保真构建
    高保真构建

可复制，可验证，已验证的构建

    一个可复制的构建, 当重复构建时, 会产生相同的结果
    一个可验证的构建, 记录足够的信息以精确地描述如何重复它
        goversion 可以获取 Go 的二进制文件中包含的一个表示它们的 Go 版本的字符串
        安装 go get -u rsc.io/goversion
        校验 goversion goapp
        如果是使用 vgo 构建的程序，使用 goversion -m vgoapp，可以查到该程序依赖包的信息
    一个已验证的构建, 可以检查是否使用了预期的源代码

[参考](https://lingchao.xin/post/vgo-tour.html)
