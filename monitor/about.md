# 服务器监控

0 监控什么

    请求次数（QPS），最大响应时间，平均响应时间
    程序占用资源情况，如内存使用情况

[infoq 翻译自 Baron 的博客 1](http://www.infoq.com/cn/articles/chat-monitoring-part01)

1 prometheus + grafana

    安装
        go get -u -v github.com/prometheus/prometheus/cmd/...
    编写配置文件 prometheus.yml (注意缩进)
        global:
        scrape_interval:     15s
        evaluation_interval: 30s

        scrape_configs:
        - job_name: prometheus

        static_configs:
            - targets: ['localhost:9090']
    启动
        prometheus --web.listen-address="0.0.0.0:9090"
    会生成 data/wal 目录，里面保存预写式日志

[参考 1](http://cjting.me/linux/use-prometheus-to-monitor-server/)
[参考 2](https://www.hi-linux.com/posts/25047.html)

2 Elastic Stack

    Metricbeat 是一个专门用来获取服务器或应用服务内部运行指标数据的收集程序

[参考](https://gocn.io/article/285)
