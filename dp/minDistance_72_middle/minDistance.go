package main

/**
 * @Author: chenpengfei
 * @Date: 2024/2/3 下午6:24
 * @Desc: 编辑距离
给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符

*/

// minDistance
// 定义dp[i][j]为word1第i个位置word1[0:i+1]和word2第j个位置word2[0:j+1]的最小编辑距离
// 转移方程：
// if (word1.charAt(i - 1) == word2.charAt(j - 1)) {
//   - dp[i][j] = dp[i - 1][j - 1];
//   - }
//   - else {
//   - dp[i][j] = Math.min(Math.min(dp[i - 1][j - 1], dp[i][j - 1]), dp[i - 1][j]) + 1;
//   - }
//     @Description:
//     @param word1
//     @param word2
//     @return int
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}
	var dp = make([][]int, m+1)
	// 初始化dp数组
	// j = 0说明是空字符串
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
		if i > 0 {
			dp[i][0] = dp[i-1][0] + 1
		}
	}
	// i = 0说明word1是空字符串
	for j := 1; j < n+1; j++ {
		dp[0][j] = dp[0][j-1] + 1
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = minTwo(minTwo(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[m][n]
}

func minTwo(a, b int) int {
	if a < b {
		return a
	}
	return b
}
