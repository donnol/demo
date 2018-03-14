# MD5 消息摘要算法

[一种被广泛使用的密码散列函数，可以产生出一个 128 位（16 字节）的散列值（hash value），用于确保信息传输完整一致](https://zh.wikipedia.org/wiki/MD5)

    MD5是输入不定长度信息，输出固定长度128-bits的算法
    经过程序流程，生成四个32位数据，最后联合起来成为一个128-bits散列。基本方式为，求余、取余、调整长度、与链接变量进行循环运算，得出结果
    一般会加上一组随机字符串（Random string），防止有人利用碰撞（Collision）解密字符串
    1996年后被证实存在弱点，可以被加以破解，对于需要高度安全性的数据，专家一般建议改用其他算法