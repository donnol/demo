# elf(Executable and Linkable Format) 可执行和链接格式

[linux 中的目标文件](https://ctf-wiki.github.io/ctf-wiki/executable/elf/elf_structure/)

主要有以下类型

    可重定位文件（Relocatable File），包含由编译器生成的代码以及数据
        链接器会将它与其它目标文件链接起来从而创建可执行文件或者共享目标文件。在 Linux 系统中，这种文件的后缀一般为 .o
    可执行文件（Executable File），就是我们通常在 Linux 中执行的程序
    共享目标文件（Shared Object File），包含代码和数据，这种文件是我们所称的库文件，一般以 .so 结尾。一般情况下，它有以下两种使用情景
        链接器（Link eDitor, ld）可能会处理它和其它可重定位文件以及共享目标文件，生成另外一个目标文件。
        动态链接器（Dynamic Linker）将它与可执行文件以及其它共享目标组合在一起生成进程镜像。

目标文件由汇编器和链接器创建，是文本程序的二进制形式，可以直接在处理器上运行。那些需要虚拟机才能够执行的程序(Java)不属于这一范围。
