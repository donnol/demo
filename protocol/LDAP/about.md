# LDAP

[轻型目录存取协定(Lightweight Directory Access Protocol)](https://zh.wikipedia.org/wiki/%E8%BD%BB%E5%9E%8B%E7%9B%AE%E5%BD%95%E8%AE%BF%E9%97%AE%E5%8D%8F%E8%AE%AE)

介绍

    通过IP协议提供访问控制和维护分布式信息的目录信息
    用于开发内部网和与互联网程序共享用户、系统、网络、服务和应用服务
    LDAP的一个常用用途是单点登录，用户可以在多个服务中使用同一个密码，通常用于公司内部网站的登录中
    在TCP/IP之上定义了一个相对简单的升级和搜索目录的协议

协议内容

    LDAP 目录的条目(entry)由属性(attribute)的一个聚集组成，并由一个唯一性的名字引用，即专有名称(distinguished name，DN)
    LDAP 目录与普通数据库的主要不同之处在于数据的组织方式，它是一种有层次的、树形结构
    LDAP 目录条目可描述一个层次结构，这个结构可以反映一个政治、地理或者组织的范畴

[特点](https://segmentfault.com/a/1190000002607140)

    树结构，查询快，写入慢
