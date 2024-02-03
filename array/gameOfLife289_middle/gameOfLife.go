package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2022/9/11 4:38 PM
 * @Desc: 根据百度百科，生命游戏，简称为 生命 ，是英国数学家约翰·何顿·康威在 1970 年发明的细胞自动机。

给定一个包含 m × n个格子的面板，每一个格子都可以看成是一个细胞。每个细胞都具有一个初始状态： 1 即为 活细胞 （live），或 0 即为 死细胞 （dead）。每个细胞与其八个相邻位置（水平，垂直，对角线）的细胞都遵循以下四条生存定律：

如果活细胞周围八个位置的活细胞数少于两个，则该位置活细胞死亡；
如果活细胞周围八个位置有两个或三个活细胞，则该位置活细胞仍然存活；
如果活细胞周围八个位置有超过三个活细胞，则该位置活细胞死亡；
如果死细胞周围正好有三个活细胞，则该位置死细胞复活；
下一个状态是通过将上述规则同时应用于当前状态下的每个细胞所形成的，其中细胞的出生和死亡是同时发生的。给你 m x n 网格面板 board 的当前状态，返回下一个状态。
*/

func main() {
	var b = [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	gameOfLife(b)
	fmt.Println(b)
}

//
// gameOfLife
//  @Description: 循环遍历每一点，将满足上数规律的点设置为死/活细胞
// 使用位运算来进行细胞标记，不能直接将值设置为最终值，会影响其他细胞的条件判断
//  @param board
//
func gameOfLife(board [][]int) {
	direction := [][]int{{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			cnt := 0
			for _, dir := range direction {
				if (i+dir[0]) >= 0 && (i+dir[0]) < len(board) && (j+dir[1]) >= 0 && (j+dir[1]) < len(board[0]) {
					// 有效的细胞位置
					if board[i+dir[0]][j+dir[1]]&0b1 == 1 {
						cnt++
					}
				}
			}
			if board[i][j]&0b1 == 1 {
				// 如果是活细胞
				if cnt < 2 || cnt > 3 {
					// 变成死细胞
					board[i][j] |= 0b10
				}
			}
			if board[i][j]&0b1 == 0 {
				// 如果是死细胞
				if cnt == 3 {
					// 变成活细胞
					board[i][j] |= 0b10
				}
			}
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 0b11 {
				board[i][j] = 0
			}
			if board[i][j] == 0b10 {
				board[i][j] = 1
			}
		}
	}
}
