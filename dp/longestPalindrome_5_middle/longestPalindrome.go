package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2024/2/3 下午5:06
 * @Desc: 给你一个字符串 s，找到 s 中最长的回文子串。

如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。
*/

func main() {
	s := "aaaa"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	// 定义dp[i][j]为s[i:j+1]是否为回文串
	// dp[i][j]是否为回文串取决于s[i]，s[j]和dp[i+1][j-1]
	// 并且i<=j必须成立，可以减少循环次数
	begin, maxLen := 0, 1
	var dp = make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
	// 一定要用j作为外层循环，避免出现遍历到dp[i][j]时
	// dp[i+1][j-1]还没遍历到
	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ {
			if i == j {
				dp[i][j] = true
			} else {
				if s[i] == s[j] {
					if j-i == 1 {
						// 相邻情况
						dp[i][j] = true
					} else {
						dp[i][j] = dp[i+1][j-1]
					}
				}
			}
			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+maxLen]
}
