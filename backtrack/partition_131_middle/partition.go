package main

import (
	"fmt"
	"slices"
)

/**
 * @Author: chenpengfei
 * @Date: 2024/2/3 下午3:16
 * @Desc: 给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。

回文串 是正着读和反着读都一样的字符串。
*/

func main() {
	s := "aab"
	r := partition2(s)
	fmt.Println(r)
}

// partition
// 思路: 使用回溯算法, 其中track为已经都是回文的子串，choose为剩下的字符串，
//   - 每次从剩下的字符串中截取一个left，如果left是回文则添加进入track中，否则left向右延伸一位继续判断
//   - 做法和第93题一样，算法优化的方案，修改isValid函数，可以利用一个dp数组来判断一个字符串是否为回文
//     @Description:
//     @param s
//     @return [][]string
func partition(s string) [][]string {
	var r [][]string
	trackBack(0, nil, s, &r)
	return r
}

func trackBack(start int, track []string, s string, result *[][]string) {
	if start == len(s) {
		var copyTrack = make([]string, len(track))
		copy(copyTrack, track)
		*result = append(*result, copyTrack)
	}
	for i := start; i < len(s); i++ {
		left := s[start : i+1]
		if isValid(left) {
			track = append(track, left)
			trackBack(i+1, track, s, result)
			track = track[:len(track)-1]
		}
	}
}

func isValid(s string) bool {
	if len(s) == 1 {
		return true
	}
	bs := []byte(s)
	slices.Reverse(bs)
	return s == string(bs)
}

func partition2(s string) [][]string {
	// 先生成一个dp数组用来表示是否回文
	// 初始化把所有为都置为true
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		for j := 0; j < len(s); j++ {
			dp[i][j] = true
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
		}
	}

	var path []string
	result := new([][]string)
	helper(dp, s, 0, path, result)
	return *result
}

func helper(dp [][]bool, s string, index int, path []string, res *[][]string) {
	if index == len(s) {
		c := make([]string, len(path))
		copy(c, path)
		*res = append(*res, c)
		return
	}

	for i := index; i < len(s); i++ {
		if dp[index][i] {
			path = append(path, s[index:i+1])
			helper(dp, s, i+1, path, res)
			path = path[:len(path)-1]
		}
	}
}
