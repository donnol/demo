# README

1 执行步骤

    先执行 gen_cert.sh 生成证书
        如果运行脚本时出现 ./gen_cert.sh: line 2: $'\r': command not found
        先运行 sed -i 's/\r$//' gen_cert.sh
        然后再运行脚本
    启动服务器
        go build
        ./server
