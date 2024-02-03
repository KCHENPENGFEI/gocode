package main

import (
	"testing"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/8/27 上午10:54
 * @Desc:
 */

var nums = []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}

func BenchmarkLongestConsecutive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		longestConsecutive(nums)
	}
}

func BenchmarkLongestConsecutive1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		longestConsecutive(nums)
	}
}
