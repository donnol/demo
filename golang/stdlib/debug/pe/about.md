# pe(Portable Executable) 可移植性可执行文件

[一种用于可执行文件、目标文件和动态链接库的文件格式，主要使用在 32 位和 64 位的 Windows 操作系统上](https://zh.wikipedia.org/wiki/%E5%8F%AF%E7%A7%BB%E6%A4%8D%E5%8F%AF%E6%89%A7%E8%A1%8C)

封装了以下信息

    动态链接库、API导入和导出表、资源管理数据和线程局部存储数据

节（section）是 PE 文件中存储数据的划分。通常，加载到内存后，同一节的数据具有相同的内存访问属性（可读/可写/可执行等）
