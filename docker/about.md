# docker

关键字：标准化

码头工人，集装箱

一般地开发流程是，开发做完开发，将程序和需要的环境给运维，运维据此部署到服务器上。中间可能出现一个问题，开发环境和服务器环境不一致。如果有多个服务器在跑，同时部署和维护这么多环境就会消耗大量时间和精力，是一个不小的成本。

## 基础命令(以 golang 镜像为例)

    镜像相关：
        # 查看本机已有镜像
        sudo docker image ls
        # golang 是镜像名，当镜像不存在时，会自动下载
        sudo docker run golang
    容器相关：
        # 运行 go version 命令
        sudo docker run golang go version
        # 查看所有容器，没有 -a 选项列出的是正在运行的容器
        sudo docker ps -a
        # 查看容器的环境变量
        sudo docker run golang env

## docker file

在目录里编译好程序，生成命令 nm ，然后编写 dockerfile

```docker
# 编写一个运行 Go 程序输出 Hello World 的镜像
FROM golang
COPY nm WORKDIR/
RUN WORKDIR/nm # nm 是命令名
```

生成镜像 sudo docker build -t mydocker . # -t 指定镜像名和标签，不写标签，默认是 latest
运行容器 sudo docker run mydocker WORKDIR/nm

## 和 Go 配合

Go 本身在开发完成后编译生成一个无依赖的二进制可执行文件，哪怕是开发环境和运行环境不一致，也可以通过自身工具交叉编译到目标环境，那是否就可以说 Go 开发中就不需要使用到 docker 了呢？

以下有几种情况，使用 docker 会是更好选择：

    1 程序运行时，除了本身的二进制文件外，还需要用到其它文件，如配置文件，静态文件等
        当然，静态文件也可以通过一些库将它们打包进二进制文件中
    2 服务器数量多，逐个交叉编译到目标系统的时间也是很大消耗
    3 如果将来打算将开发语言更换，那么打包进docker会是更好选择
    4 资源隔离和限制，当单个程序并不允许它占用太多资源时
    5 多版本测试，多配置测试，如数据库等第三方中间件

## SSh

如果容器里面有安装 openssh 的话，配置好端口映射后，就可以通过 ssh 命令登录进容器了。但是，很多时候，我们只是在容器里运行一个服务，容器里并没有安装 openssh。这时候，可以通过 docker 提供的命令来执行一些命令(命令必须是容器文件系统里存在的)，如 docker exec -it [mycontainer] cmd
