package main

import (
	"fmt"
	"strings"
)

type NQueen struct {
	res [][]string
}

func solveNQueens(n int) [][]string {
	q := NQueen{
		res: [][]string{},
	}
	board := makeBoard(n)
	q.backtrack(board, 0)
	return q.res
}

func makeBoard(n int) []string {
	board := make([]string, n)
	rows := make([]string, n)
	for i := range rows {
		rows[i] = "."
	}
	row := strings.Join(rows, "")
	for i := range board {
		board[i] = row
	}
	return board
}

func copyStrArray(board []string) []string {
	temp := make([]string, len(board))
	for i, v := range board {
		temp[i] = v
	}
	return temp
}

func (nq *NQueen) backtrack(board []string, row int) {
	n := cap(board)
	// 触发结束条件
	if row == n {
		fmt.Printf("获得答案: %v\n", board)
		nq.res = append(nq.res, copyStrArray(board))
		return
	}
	for col := 0; col < n; col++ {
		// 排除不合法选择
		if !isValid(board, row, col) {
			continue
		}
		// 做选择
		SetChar(&board[row], col, "Q")
		fmt.Printf("第%d行，第%d列，做选择: %v\n", row, col, board)
		// 进入下一行决策
		nq.backtrack(board, row+1)
		// 撤销选择
		SetChar(&board[row], col, ".")
		fmt.Printf("第%d行，第%d列，撤销选择: %v\n", row, col, board)
	}
}

func isValid(board []string, row, col int) bool {
	n := cap(board)
	// 检查列是否有皇后互相冲突
	for i := 0; i < row; i++ {
		if i < len(board) && col < len(board[i]) && board[i][col] == 'Q' {
			return false
		}
	}
	// 检查右上方是否有皇后互相冲突
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if i < len(board) && j < len(board[i]) && board[i][j] == 'Q' {
			return false
		}
	}
	// 检查左上方是否有皇后互相冲突
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if i < len(board) && j < len(board[i]) && board[i][j] == 'Q' {
			return false
		}
	}
	return true
}

// SetChar 修改字符串对应字节
func SetChar(targetStr *string, pos int, newChar string) {
	if len(*targetStr) <= pos {
		for i := pos - len(*targetStr); i >= 0; i-- {
			*targetStr += "."
		}
	}
	temp := []rune(*targetStr)
	temp[pos] = []rune(newChar)[0]
	*targetStr = string(temp)
}

func main() {
	res := solveNQueens(4)
	for _, r := range res {
		for _, l := range r {
			fmt.Println(l)
		}
		fmt.Println("--------------------------------")
	}

}
