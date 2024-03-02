package main

/**
 * @Author: chenpengfei
 * @Date: 2023/12/3 下午2:44
 * @Desc: 给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。
	维护两个map保存行有0和列有0的情况
*/

func setZeroes(matrix [][]int) {
	// 两个set保存需要置零的行和列
	rowMap := make(map[int]struct{}, len(matrix))
	colMap := make(map[int]struct{}, len(matrix[0]))

	for i, row := range matrix {
		for j, num := range row {
			if num == 0 {
				rowMap[i] = struct{}{}
				colMap[j] = struct{}{}
			}
		}
	}
	for i, row := range matrix {
		_, ok1 := rowMap[i]
		for j := range row {
			if ok1 {
				matrix[i][j] = 0
			}
			_, ok2 := colMap[j]
			if ok2 {
				matrix[i][j] = 0
			}
		}
	}
}
