package main

/**
 * @Author: chenpengfei
 * @Date: 2023/12/3 下午2:44
 * @Desc: 给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。
	使用中间态-1来临时存储被修改为0的位置，并且维护两个map来保存行中有0或者列中有0的情况
	空间复杂度为O(m + n)
*/

func setZeroes(matrix [][]int) {
	// 矩阵有效
	m := len(matrix)
	n := len(matrix[0])
	lineZero := make(map[int]struct{}, m)
	colZero := make(map[int]struct{}, n)
	for i := 0; i < m; i++ {
		lineHaveZero := false
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 && !lineHaveZero {
				// 第一次遇到0，将本行和本列全部设置为0
				for k := 0; k < n; k++ {
					if matrix[i][k] != 0 {
						// 中间态
						matrix[i][k] = -1
					}
				}
				lineHaveZero = true
				// 将本列全部设置为0
				for k := 0; k < m; k++ {
					if matrix[k][j] != 0 {
						matrix[k][j] = -1
					}
				}
				lineZero[i] = struct{}{}
				colZero[j] = struct{}{}
			} else if matrix[i][j] == 0 {
				// 不是第一遇到0，只需要将本列全部设置为0
				for k := 0; k < m; k++ {
					if matrix[k][j] != 0 {
						matrix[k][j] = -1
					}
				}
				colZero[j] = struct{}{}
			}
		}
	}
	// 遍历矩阵将-1全部置换成0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			_, ok1 := lineZero[i]
			_, ok2 := colZero[j]
			if matrix[i][j] == -1 && (ok1 || ok2) {
				matrix[i][j] = 0
			}
		}
	}
}
