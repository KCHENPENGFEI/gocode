package main

/**
 * @Author: chenpengfei
 * @Date: 2024/2/3 下午5:46
 * @Desc: 最长公共子序列
给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。

一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。

例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。
*/

// longestCommonSubsequence
//
//	@Description: * 做法：定义dp[i][j]为s1[0...i]和s2[0...j]的最长公共子串的长度
//
// * 那么有以下转移方程
// * if (s1[i] == s2[j]) {
// *     dp[i][j] = dp[i - 1][j - 1] + 1;
// * }
// * else {
// *     dp[i][j] = Math.max(dp[i - 1][j], dp[i][j - 1]);
// * }
// *
//
//	@param text1
//	@param text2
//	@return int
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	var dp = make([][]int, m)
	// 初始化dp数组(处理i=0和j=0的情况)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		if text1[i] == text2[0] {
			dp[i][0] = 1
		} else {
			if i == 0 {
				dp[i][0] = 0
			} else {
				dp[i][0] = dp[i-1][0]
			}
		}
	}
	// j=0
	for j := 1; j < n; j++ {
		if text2[j] == text1[0] {
			dp[0][j] = 1
		} else {
			dp[0][j] = dp[0][j-1]
		}
	}
	// dp转移方程
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if text1[i] == text2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = maxTwo(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m-1][n-1]
}

func maxTwo(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// longestCommonSubsequence1
//
//	@Description: 可以定义dp[i][j]是s1[0...i-1]和s2[0...j-1]的最长公共子串的长度
//	这样可以减少初始化过程，此时i和j都是从1开始
//	@param text1
//	@param text2
//	@return int
func longestCommonSubsequence1(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	var dp = make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = maxTwo(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}
