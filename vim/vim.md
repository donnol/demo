# vim

> h,j,k,l    上，下，左，右
>
> ctrl-f     上翻一页
>
> ctrl-b     下翻一页
>
>%          跳到与当前括号匹配的括号处，如当前在{，则跳转到与之匹配的}处
>
>w     跳到下一个字首，按标点或单词分割
>
>W     跳到下一个字首，长跳，如end-of-line被认为是一个字
>
>e     跳到下一个字尾
>
>E     跳到下一个字尾，长跳
>
>b     跳到上一个字
>
>B     跳到上一个字，长跳
>
>0     跳至行首，不管有无缩进，就是跳到第0个字符
>
>^     跳至行首的第一个字符
>
>$     跳至行尾
>
>gg     跳至文件的第一行
>
>gd     跳至当前光标所在的变量的声明处
>
>[N]G     跳到第N行，如0G，就等价于gg，100G就是第100行
>
>fx     在当前行中找x字符，找到了就跳转至
>
>;     重复上一个f命令，而不用重复的输入fx
>
>tx     与fx类似，但是只是跳转到x的前一个字符处
>
>Fx     跟fx的方向相反
>
>),(     跳转到上/下一个语句
>
>\*     查找光标所在处的单词，向下查找
>
>\#     查找光标所在处的单词，向上查找
>
>`.     跳转至上次编辑位置
>
>在屏幕上移动
>
>H     移动光标到当前屏幕上最上边的一行
>
>M     移动光标到当前屏幕上中间的一行
>
>L     移动光标到当前屏幕上最下边的一行
>
>书签
>
>ma     把当前位置存成标签a
>
>`a     跳转到标签a处
>
>编辑
>r     替换一个字符
>
>J     将下一行和当前行连接为一行
>
>cc     删除当前行并进入编辑模式
>
>cw     删除当前字，并进入编辑模式
>
>c$     擦除从当前位置至行末的内容，并进入编辑模式
>
>s     删除当前字符并进入编辑模式
>
>S     删除光标所在行并进入编辑模式
>
>xp     交换当前字符和下一个字符
>
>u     撤销
>
>ctrl+r     重做
>
>.     重复上一个编辑命令
>
>~     切换大小写，当前字符
>
>g~iw     切换当前字的大小写
>
>gUiw     将当前字变成大写
>
>guiw     将当前字变成小写
>
> \>\>     将当前行右移一个单位
>
> <<     将当前行左移一个单位(一个tab符)
>
>==     自动缩进当前行
>