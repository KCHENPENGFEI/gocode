package main

/**
 * @Author: chenpengfei
 * @Date: 2024/2/3 下午4:37
 * @Desc:
 */
func uniquePaths(m int, n int) int {
	var dp = make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
