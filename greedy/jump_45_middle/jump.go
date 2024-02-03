package main

/**
 * @Author: chenpengfei
 * @Date: 2024/1/27 下午4:33
 * @Desc:
 */

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	step := 0
	help(nums, &step)
	return step
}

func help(nums []int, depth *int) {
	*depth += 1
	// 从最后一个元素往前找
	minIdx := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i]+i >= len(nums)-1 {
			if i == 0 {
				return
			}
			minIdx = i
		}
	}
	nums = nums[:minIdx+1]
	help(nums, depth)
}
