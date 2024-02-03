package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2023/12/30 下午3:19
 * @Desc:
 */

func main() {
	b := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	w := "ABCCEP"
	fmt.Println(exist(b, w))
}

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			var newBoard [][]byte
			for k := 0; k < len(board); k++ {
				var t = make([]byte, len(board[0]))
				copy(t, board[k])
				newBoard = append(newBoard, t)
			}
			if newBoard[i][j] == word[0] {
				if dfs(i, j, newBoard, word, 0) {
					return true
				}
			}
		}
	}
	return false
}

func dfs(i, j int, board [][]byte, word string, index int) bool {
	// 优先处理不满足的情况
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
		return false
	}
	if board[i][j] != word[index] {
		return false
	}
	if len(word) == index+1 {
		return true
	}
	c := board[i][j]
	board[i][j] = '#'
	if dfs(i-1, j, board, word, index+1) || dfs(i+1, j, board, word, index+1) ||
		dfs(i, j-1, board, word, index+1) || dfs(i, j+1, board, word, index+1) {
		return true
	}
	board[i][j] = c
	return false
}
