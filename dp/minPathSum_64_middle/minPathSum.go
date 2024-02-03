package main

/**
 * @Author: chenpengfei
 * @Date: 2024/2/3 下午4:44
 * @Desc:给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。
*/

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var dp = make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
			} else if i == 0 {
				dp[i][j] = grid[i][j] + dp[i][j-1]
			} else if j == 0 {
				dp[i][j] = grid[i][j] + dp[i-1][j]
			} else {
				dp[i][j] = minTwo(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}
	return dp[m-1][n-1]
}

func minTwo(a, b int) int {
	if a < b {
		return a
	}
	return b
}
