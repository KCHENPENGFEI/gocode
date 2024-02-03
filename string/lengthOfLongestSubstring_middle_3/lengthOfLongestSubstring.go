package main

import (
	"fmt"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/9/8 上午10:50
 * @Desc: 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
	方法1：针对每个下标i，找到以i为起点的最长不重复子串
	方法2：采用滑动串口法，求取连续子串问题就用滑动串口法
*/

func main() {
	s := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	for i := 0; i < len(s); i++ {
		curLen := findSubString(i, s)
		if curLen > maxLen {
			maxLen = curLen
		}
	}
	return maxLen
}

func findSubString(start int, s string) int {
	var arr [128]int8
	for i := start; i < len(s); i++ {
		if arr[int8(s[i])] == 1 {
			return i - start
		} else {
			arr[int8(s[i])] = 1
		}
	}
	return len(s) - start
}
