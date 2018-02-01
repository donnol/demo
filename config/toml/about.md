# toml

一种可以无歧义地映射为哈希表的配置文件格式，非常简洁、优雅，作者是 GitHub 　联合创始人　 Tom Preston-Werner，要求使用 UTF-8 编码

```toml
# 这是注释

str = "i am a string" # 这是一个字符串

escape = "\n\b\t" # 转义

int = 1 # 整型
negInt = -1 # 负整型

float = 0.1 # 浮点
negFloat = -0.1 # 负浮点

boolTrue = true # 布尔 真
boolFalse = false # 布尔 假

date = 1979-05-27T07:32:00Z # 时间，ISO 8601　完整格式

# 数组不可以混合类型，多维数据里面的任意最小数组必须是相同类型
array1 = [1, 2, 3] # 数组
array2 = [[1, 2, 3], ["a", "b", "c"]] # 这样可以
array3 = [1, "b", "c"] # 这样不行
array4 = [
    1,
    2,
    3,
]

# 表
[table]
    key1 = "value1"
    key2 = 2
[table.innerTable] # 嵌入
    innerKey1 = "innerValue1"
    innerKey2 = 2
[a.b.c.d] # 事先不声明父表
    noDesignTableKey1 = 1
# 不能多次定义键和表格
[e]
f = 1
[e] # 不能再次定义
f = 1
[e.f] # 不能再次定义
g = 1

# 表格数组
[[tableArray]]
    key1 = 1
[[tableArray]]
    key1 = 2
    [[tableArray.innerTableArray]] # 嵌套
        key2 = 3
    [tableArray.innerTableArray] # 已有定义，这里会出错
        key3 = 4
```

## 校验工具

    go get -u -v github.com/BurntSushi/toml/cmd/tomlv
    tomlv test.toml
