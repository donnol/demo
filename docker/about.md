# docker

关键字：标准化

码头工人，集装箱

Docker 镜像(Image)，就相当于是一个 root 文件系统

    Docker 镜像是一个特殊的文件系统，除了提供容器运行时所需的程序、库、资源、配置等文件外，还包含了一些为运行时准备的一些配置参数（如匿名卷、环境变量、用户等）。镜像不包含任何动态数据，其内容在构建之后也不会被改变。
    每个镜像都由很多层次构成，Docker 使用 Union FS 将这些不同的层结合到一个镜像中去

容器

    以镜像启动，在运行过程中的修改，可以保存为新的镜像

数据管理

    数据卷 sudo docker volume create [volume_nam]
        不随容器删除而删除
    挂载主机目录
        测试时使用

网络管理

    端口 -p host:port:containerhost:containerport/[tcp|udp]
    容器互联
        # 新建网络
        sudo docker network create -d bridge [my-net]
        # 容器运行时使用指定网络
        sudo docker run --network [my-net] image
    DNS配置
        1  可以在文件 /etc/docker/daemon.json 中添加以下内容
            {"dns" : ["114.114.114.114","8.8.8.8"] }
        2 在容器启动时指定 --hostname=hostname --dns=IP_ADDRESS
    高级
        当 Docker 启动时，会自动在主机上创建一个  docker0  虚拟网桥，实际上是 Linux 的一个 bridge，可以理解为一个软件交换机。它会在挂载到它的网口之间进行转发。
        当创建一个 Docker 容器的时候，同时会创建了一对  veth pair  接口（当数据包发送到一个 接口时，另外一个接口也可以收到相同的数据包）。这对接口一端在容器内，即  eth0 ；另 一端在本地并被挂载到  docker0  网桥，名称以  veth  开头（例如  vethAQI2QT ）。通过这种 方式，主机可以跟容器通信，容器之间也可以相互通信。Docker 就创建了在主机和所有容器 之间一个虚拟共享网络

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
        # 清除处于终止状态的容器
        sudo docker container prune

## docker file

在目录里编译好程序，生成命令 nm ，然后编写 dockerfile

1 scratch 作为基础镜像

```docker
FROM scratch # 使用 scratch 作为基础镜像时，意味着你不以任何镜像为基础，接下来所写的指令将作为镜像第一层开始存在
COPY nm /
CMD ["/nm"]
```

生成镜像 sudo docker build -t mydocker2 -f dockerfile2 . # -f 指定 dockerfile
运行容器 sudo docker run mydocker2

2 以 golang 作为基础镜像

```docker
# 编写一个运行 Go 程序输出 Hello World 的镜像
FROM golang
COPY nm WORKDIR/
# nm 是命令名，RUN 命令需要镜像里存在 /bin/sh 才可以使用，如果没有会报错：
#   OCI runtime create failed: container_linux.go:295: starting container process caused
#   "exec:\"/bin/sh\": stat /bin/sh: no such file or directory": unknown
RUN WORKDIR/nm
```

生成镜像 sudo docker build -t mydocker . # -t 指定镜像名和标签，不写标签，默认是 latest
运行容器 sudo docker run mydocker WORKDIR/nm

3 一个简单服务器的镜像

```docker
FROM ubuntu
COPY server /
CMD ["/server"] # 容器运行时的默认命令
```

生成镜像 sudo docker build -t myserver .
运行容器 sudo docker run -p 8520:8520 -d myserver # -p 将容器的端口映射到虚拟机端口; -d 后台运行，使用 sudo docker logs CONTAINER_ID 查看输出

    出现错误: standard_init_linux.go:195: exec user process caused "no such file or directory"
    分析：因为网络服务需要DNS解析器，而 scratch 镜像是什么都没有的，所以报错无法找到文件
    解决：
        1 换另外的基础镜像，如 busybox,alpine,golang:alpine,ubuntu 等
        其中，只有换用 ubuntu 镜像后成功了，其它一样还是不行

[参考](http://geek.csdn.net/news/detail/248772)

## 发送镜像

1 通过仓库

    登录 sudo docker login
    拉 sudo docker pull [image]
    推 sudo docker push user/image

2 通过文件

    保存成文件 sudo docker save -o <save image to path> <image name>
    将文件从虚拟机保存到本地
        # 使用 sftp 在主机建立与虚拟机的连接
        sftp -P 3022 jd@127.0.0.1  # 我这里是虚拟机发送到本机，端口设置参考 vm/win10_vb.md
        # 获取文件
        get <path to image tar file>
    从文件加载 sudo docker load -i <path to image tar file>

## 和 Go 配合

Go 本身在开发完成后编译生成一个无依赖的二进制可执行文件，哪怕是开发环境和运行环境不一致，也可以通过自身工具交叉编译到目标环境，那是否就可以说 Go 开发中就不需要使用到 docker 了呢？

以下有几种情况，使用 docker 会是更好选择：

    1 程序运行时，除了本身的二进制文件外，还需要用到其它文件，如配置文件，静态文件等
        当然，静态文件也可以通过一些库将它们打包进二进制文件中
    2 服务器数量多，逐个交叉编译到目标系统的时间也是很大消耗
    3 如果将来打算将开发语言更换，那么打包进docker会是更好选择
    4 资源隔离和限制，当单个程序并不允许它占用太多资源时
    5 多版本测试，多配置测试，如数据库等第三方中间件

## SSH

如果容器里面有安装 openssh 的话，配置好端口映射后，就可以通过 ssh 命令登录进容器了。但是，很多时候，我们只是在容器里运行一个服务，容器里并没有安装 openssh。这时候，可以通过 docker 提供的命令来执行一些命令(命令必须是容器文件系统里存在的)，如 docker exec -it [mycontainer] cmd

[参考](https://www.gitbook.com/book/yeasy/docker_practice/details)
