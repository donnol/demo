# 架构

1 三层架构

    界面层(User Interface layer)
    业务层(Business Logic Layer)
    数据层(Data access layer)

    与 MVC
        联系:
            - 初看起来比较像，容易搞混(如果这也算联系的话);
            - MVC 模式可以用于3-tiers 架构的展示层
        区别:
            1 关注的重点不同:
                MVC
                    关注的重点在于表现层的代码组织方式，通过降低代码间的耦合度，使代码更易维护
                3-tires
                    关注系统的分布，便于提升系统性能，增加系统功能
            2 拓扑结构不同:
                MVC
                    是可以三角结构，视图向控制器发送更新，控制器更新模型，视图可以直接从模型更新
                3-tires
                    一定是是线性结构，展示层即客户端不能直接与数据层通信，也就是说客户端展示层与数据层的通信必须经过中间层即业务层

[知乎](https://www.zhihu.com/question/21851341/answer/251629127)
