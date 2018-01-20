# 自建 git 服务

1 在服务器上

    安装 git
    sudo apt install git

    添加用户 git
    sudo useradd git
    设置用户密码
    sudo passwd git
    修改用户 shell
    vim /etc/passwd
    /bin/bash -> /usr/bin/git-shell
    创建用户目录
    sudo mkdir /home/git
    并在/home/git 目录里创建 .ssh 目录
    在 .ssh 目录里新建文件 authorized_keys
    将 /home/git 目录转给 git
    sudo chown -R git:git /home/git

    到任意目录 创建仓库
    sudo git init --bare shark.git
    并将拥有者改为 git
    sudo chown -R git:git shark.git
    否则客户机将无法获取/更新仓库

2 在本地机上

    修改 /etc/ssh/ssh_config
    将 port 的值改为服务器 /etc/ssh/sshd_config 文件里暴露的 port 值
    将本地的 ~/.ssh/id_rsa.pub 文件里的内容复制到服务器的 authorized_keys 文件里

    获取仓库
    git clone git@[server]:/path/to/yourrepository.git
