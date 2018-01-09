package a

// A A 类型
type A string

// B B 类型`当使用别名时，无法被 godef 找到定义，使用 gogetdoc 可以
type B = string
