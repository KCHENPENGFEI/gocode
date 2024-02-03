package main

/**
 * @Author: chenpengfei
 * @Date: 2024/1/28 下午2:06
 * @Desc:
 */

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	// 分成两个队列
	// 第一个不包含最后一个元素
	// 第二个不包含第一个元素
	return maxTwo(rob1(nums[1:]), rob1(nums[:len(nums)-1]))
}

func rob1(nums []int) int {
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
	return s[len(s)-1]
}

func maxTwo(a, b int) int {
	if a > b {
		return a
	}
	return b
}
