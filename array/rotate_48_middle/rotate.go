package main

import (
	"fmt"
	"sync"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/12/10 下午4:38
 * @Desc: 给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。

你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。

*/

func main() {
	m := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotate(m)
	fmt.Println(m)
}

func rotate(matrix [][]int) {
	n := len(matrix)
	var wg sync.WaitGroup
	for i := 0; i <= n/2; i++ {
		for j := i; j < n-i-1; j++ {
			wg.Add(1)
			go func(i, j int) {
				defer wg.Done()
				rotateOne(i, j, n, matrix)
			}(i, j)
		}
	}
	wg.Wait()
}

func rotateOne(i, j int, n int, matrix [][]int) {
	var data = matrix[i][j]
	var tmp int
	for k := 0; k < 4; k++ {
		nextI := j
		nextJ := n - 1 - i
		tmp = matrix[nextI][nextJ]
		matrix[nextI][nextJ] = data
		data = tmp
		i = nextI
		j = nextJ
	}
}
