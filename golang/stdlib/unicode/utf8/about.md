# utf8

## ASCII

[控制字符](https://zh.wikipedia.org/wiki/%E6%8E%A7%E5%88%B6%E5%AD%97%E7%AC%A6)

    33个，0~31和127，位于0x00-0x1F及0x7F
    目的：
        打印和显示控制
        数据结构化
        传输控制
        零散用途

显示

    vscode
        修改以下两个配置
        "editor.renderControlCharacters": true,
        "editor.renderWhitespace": "all",
    vim
        :set invlist
        或者
        :set list # 开启
        :set nolist # 关闭
        :set listchars=tab:-> # 字符替换显示设置，这里将tab替换为->，其它不可见字符将不显示，
            默认是 eol:$
    linux
        cat -A file
