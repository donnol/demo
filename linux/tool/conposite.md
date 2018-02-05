# 组合

1 当前目录中有多少个文件(不包含'.'开头的文件/文件夹)

```shell
ls -l . | egrep -c '^-' # egrep 可用 grep 替换
```

2 当前目录中有多少个文件(包含'.'开头的文件/文件夹)

```shell
ls -la . | egrep -c '^-' # egrep 可用 grep 替换
```

3 当前目录的文件和文件夹数量(不包含'.'开头的文件/文件夹)

```shell
ls -l | wc -l # 不太准确，因为 ls -l 会有一行 total n 的输出，wc会把这行也算进去
ls -l | egrep '^[-d]' | wc -l # 正确
```

4 当前目录的文件和文件夹数量(包含'.'开头的文件/文件夹)

```shell
ls -la | wc -l # 同上
ls -la | egrep '^[-d]' | wc -l # 正确
```

5 递归计算当前目录的文件(包括'.'开头的文件/文件夹)

```shell
find . -type f | wc -l
```

6 目录和文件数(不包括'.'开头的文件/文件夹)

```shell
tree -a | tail -l
```

7 递归计算目录数量(包含'.'开头的文件夹)

```shell
find . -type d | wc -l
```

8 根据文件扩展名计数文件数量

```shell
find . -name '*.md' | wc -l
```

9 统计当前目录中的所有文件

```shell
echo *.* | wc # 中间的数字代表文件数量
```

10 统计当前目录中的所有目录

```shell
echo */ | wc
```
