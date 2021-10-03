package main

import (
	"fmt"
	"strings"
)

func main() {
	r := solveNQueens(8)
	for _, line := range r {
		fmt.Println(line)
	}
}

func solveNQueens(n int) [][]string {
	result := make([][]string, 0, n*n*2)

	// 生成空盘面
	board := make([]string, n)
	for i := 0; i < n; i++ {
		board[i] = strings.Repeat(".", n)
	}

	// 从首行开始回溯
	backtrack(board, n, 0, &result)

	return result
}

// n一直是那个n，row从0开始，递增到n时结束，遍历行
func backtrack(board []string, n int, row int, result *[][]string) {

	// 能够把所有行都尝试过的，才能得出解；有很多种情况在中间就已经失败了，无法走到这里来
	// 能走到这里，说明又得到一种解
	if row == n {
		// board需要拷贝，否则后面的回溯会抹掉盘面
		tmp := make([]string, n)
		copy(tmp, board)
		*result = append(*result, tmp)
		return
	}

	// n!
	// 遍历列
	for col := 0; col < n; col++ {
		// 验证是否可放
		if !isValid(board, row, col) {
			continue
		}

		rowBackup := board[row]

		// 定位格子填充Q
		setStringAt(&board[row], col, 'Q')
		// 本行下完，继续下一行
		backtrack(board, n, row+1, result)

		// 回溯选定列，尝试下一列
		board[row] = rowBackup
	}
}

func isValid(board []string, row, col int) bool {
	// 排除之前行的相同列
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}

	// 排除之前行左上对角线
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	// 排除之前行右上对角线
	for i, j := row-1, col+1; i >= 0 && j < len(board); i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	return true
}

// 修改索引位置上的字符
func setStringAt(s *string, i int, r rune) {
	rs := []rune(*s)
	rs[i] = r
	*s = string(rs)
}
