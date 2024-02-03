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
	r := partition(s)
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
