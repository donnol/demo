package main

import (
	"bufio"
	"os"
)

// 一行语句
type Line string

// 块
type Block struct {
	Name      string  // 块名
	Delimeter Line    // 分隔符内容
	Tables    []Table // 相关表
}

// 表
type Table struct {
	Name     string    // 库名
	Drop     Line      // 删除语句
	Create   []Line    // 创建语句
	Comments []Comment // 注释
}

// 注释
type Comment struct {
	Column  string // 字段
	Content string // 内容
}

// 表名出现的次数
var tableCountMap map[string]int

// 设置
var setings []Line

// 扩展
var extension []Line

// 函数
type function []Line

var funcs []function

// 视图
type view []Line

var views []view

// 块
var blocks []Block

func init() {
	// 读取分块信息--哪些表组成一个块，块里面的表顺序
	jsonstr := `
		{
			'block1': ['table1'],
			'block2': ['table2']
		}
	`
	// 如果表不存在则附加到后面
}

// 扫描文件
func scanSchema(file string) (blocks []Block, err error) {
	var f *os.File
	f, err = os.Open(file)
	if err != nil {
		return
	}
	defer f.Close()

	var line string
	br := bufio.NewReader(f)
	for {
		line, err = br.ReadString("\n")
		if err != nil {
			return
		}
	}
}
