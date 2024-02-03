package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2024/1/28 下午3:30
 * @Desc:
 */

func main() {
	nums := []int{1, 3, 6, 7, 9, 4, 10, 5, 6}
	r := lengthOfLIS(nums)
	fmt.Println(r)
}

func lengthOfLIS(nums []int) int {
	var dp = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			dp[i] = 1
		} else {
			for j := i - 1; j >= 0; j-- {
				if nums[j] < nums[i] {
					dp[i] = maxTwo(dp[i], dp[j]+1)
				}
			}
		}
		if dp[i] == 0 {
			dp[i] = 1
		}
	}
	m := -1
	for i := 0; i < len(dp); i++ {
		m = maxTwo(m, dp[i])
	}
	return m
}

func maxTwo(a, b int) int {
	if a > b {
		return a
	}
	return b
}
