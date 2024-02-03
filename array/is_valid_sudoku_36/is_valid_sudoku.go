package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2022/9/7 10:26 PM
 * @Desc:请你判断一个9 x 9 的数独是否有效。只需要 根据以下规则 ，验证已经填入的数字是否有效即可。

数字1-9在每一行只能出现一次。
数字1-9在每一列只能出现一次。
数字1-9在每一个以粗实线分隔的3x3宫内只能出现一次。（请参考示例图）
*/

func main() {
	board := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	fmt.Println(isValidSudoku(board))
}

//
// isValidSudoku 构造三个map用来存储位置信息(行、列、单元格)，遍历每个数字，如果在三个map能找到相同的数字则说明数独非法
// 否则将数字写入map中
//  @Description:
//  @param board
//  @return bool
//
func isValidSudoku(board [][]byte) bool {
	rowMap := make(map[int]map[byte]bool)
	colMap := make(map[int]map[byte]bool)
	gridMap := make(map[int]map[byte]bool)

	for i, line := range board {
		for j, item := range line {
			grid := (i/3)*3 + (j / 3)
			if item == '.' {
				continue
			}
			if _, ok := rowMap[i][item]; ok {
				// 找到重复
				return false
			}
			if _, ok := colMap[j][item]; ok {
				return false
			}
			if _, ok := gridMap[grid][item]; ok {
				return false
			}
			if rowMap[i] == nil {
				rowMap[i] = make(map[byte]bool)
			}
			rowMap[i][item] = true
			if colMap[j] == nil {
				colMap[j] = make(map[byte]bool)
			}
			colMap[j][item] = true
			if gridMap[grid] == nil {
				gridMap[grid] = make(map[byte]bool)
			}
			gridMap[grid][item] = true
		}
	}
	return true
}
