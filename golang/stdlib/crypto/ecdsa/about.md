# 椭圆曲线数字签名算法（ECDSA）

[使用椭圆曲线密码（ECC）对数字签名算法（DSA）的模拟。ECDSA 于 1999 年成为 ANSI 标准，并于 2000 年成为 IEEE 和 NIST 标准](https://segmentfault.com/a/1190000012288285)

    离散对数问题
    大数分解问题
    椭圆曲线离散对数问题
        没有亚指数时间的解决方法
        因此椭圆曲线密码的单位比特强度要高于其他公钥体制
        群元素由素域中的元素数 -> 有限域上的椭圆曲线上的点
        在使用较短的密钥的情况下，ECC 可以达到与 DL 系统相同的安全级别
        好处：密钥更短，运算速度更快，签名也更加短小

参考

[椭圆曲线密码学](https://zh.wikipedia.org/wiki/%E6%A4%AD%E5%9C%86%E6%9B%B2%E7%BA%BF%E5%AF%86%E7%A0%81%E5%AD%A6)