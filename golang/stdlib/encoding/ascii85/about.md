# ascii85 即 base85

[一种二进制到文本的编码格式](https://en.wikipedia.org/wiki/Ascii85)

用于 PostScript PDF Git

比 base64 更高效

    base64
        是每6个bit映射为一个ascii字符，用64个ascii字符来表示
        每3个字节编码后对应4个字符，大小增加1/3
    base85
        是每4个字节编码为一个85进制的表示，用85个ascii字符来表示
        每4个字节编码后对应5个字符，大小增加1/4

使用 85 进制是因为

    85^5 > 2^32 > 84^5
