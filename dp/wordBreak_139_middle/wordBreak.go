package main

import (
	"fmt"
	"slices"
)

/**
 * @Author: chenpengfei
 * @Date: 2024/1/28 下午3:13
 * @Desc:
 */

func main() {
	s := "catsandog"
	w := []string{"cats", "dog", "sand", "and", "cat"}
	r := wordBreak(s, w)
	fmt.Println(r)
}

func wordBreak(s string, wordDict []string) bool {
	var dp = make([]bool, len(s))
	for i := 0; i < len(s); i++ {
		if slices.Contains(wordDict, s[0:i+1]) {
			dp[i] = true
			continue
		}
		for j := 0; j < i; j++ {
			if dp[j] && slices.Contains(wordDict, s[j+1:i+1]) {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(dp)-1]
}

func wordBreak1(s string, wordDict []string) bool {
	var dp = make([]bool, len(s))
	var m = make(map[string]struct{})
	for i := 0; i < len(wordDict); i++ {
		m[wordDict[i]] = struct{}{}
	}
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[0:i+1]]; ok {
			dp[i] = true
			continue
		}
		for j := 0; j < i; j++ {
			if _, ok := m[s[j+1:i+1]]; ok && dp[j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(dp)-1]
}
