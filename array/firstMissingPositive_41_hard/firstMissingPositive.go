package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2023/10/15 下午2:21
 * @Desc:
 */

func main() {
	nums := []int{3, 4, -1, 1}
	fmt.Println(firstMissingPositive(nums))
}

func firstMissingPositive(nums []int) int {
	n := len(nums)
	i := 0
	for i < n {
		if nums[i] == i+1 || nums[i] <= 0 {
			i++
			continue
		}
		if nums[i] > n {
			nums[i] = 0
		} else {
			index := nums[i] - 1
			if nums[index] == index+1 {
				// 目标位置已经存在正确的数值
				nums[i] = 0
			}
			//old := nums[index]
			//nums[index] = nums[i]
			nums[index], nums[i] = nums[i], nums[index]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			return i + 1
		}
	}
	return n + 1
}
