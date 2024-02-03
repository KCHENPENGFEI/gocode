package main

import (
	"fmt"
	"sort"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/8/19 下午3:36
 * @Desc:
 */

func main() {
	nums := []int{2, 1, 2}
	fmt.Println(largestPerimeter(nums))
}

func largestPerimeter(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[j] < nums[i]
	})
	for i := 0; i < len(nums)-2; i++ {
		if (nums[i+1] + nums[i+2]) > nums[i] {
			return nums[i] + nums[i+1] + nums[i+2]
		}
	}
	return 0
}
