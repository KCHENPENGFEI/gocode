package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2024/1/28 下午1:41
 * @Desc:
 */

func main() {
	n := []int{6, 3, 10, 8, 2, 10, 3, 5, 10, 5, 3}
	r := rob(n)
	fmt.Println(r)
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var s = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			s[i] = nums[i]
		} else if i == 1 {
			s[i] = maxTwo(nums[i-1], nums[i])
		} else {
			s[i] = maxTwo(s[i-1], s[i-2]+nums[i])
		}
	}
	return maxTwo(s[len(s)-1], s[len(s)-2])
}

func maxTwo(a, b int) int {
	if a > b {
		return a
	}
	return b
}
