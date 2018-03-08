# Mach-O object files

[一种用于可执行文件，目标代码，动态库，内核转储的文件格式。作为 a.out 格式的替代，Mach-O 提供了更强的扩展性，并提升了符号表中信息的访问速度](https://zh.wikipedia.org/wiki/Mach-O)

Mach-O 曾经为大部分基于 Mach 核心的操作系统所使用。NeXTSTEP，Darwin 和 Mac OS X 等系统使用这种格式作为其原生可执行文件，库和目标代码的格式。而同样使用 GNU Mach 作为其微内核的 GNU Hurd 系统则使用 ELF 而非 Mach-O 作为其标准的二进制文件格式

每个 Mach-O 文件包括一个 Mach-O 头，然后是一系列的载入命令，再是一个或多个块，每个块包括 0 到 255 个段。Mach-O 使用 REL 再定位格式控制对符号的引用。Mach-O 在两级命名空间中将每个符号编码成“对象-符号名”对，在查找符号时则采用线性搜索法
