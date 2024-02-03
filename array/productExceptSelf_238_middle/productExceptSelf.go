package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2023/10/15 下午1:42
 * @Desc: 给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。

题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。

请 不要使用除法，且在 O(n) 时间复杂度内完成此题。
*/

func main() {
	nums := []int{-1, 1, 0, -3, 3}
	fmt.Println(productExceptSelf(nums))
}

// 构造数组L[i]和R[i]来表示左乘积和右乘积，最后将L和R的数值进行相乘
func productExceptSelf(nums []int) []int {
	l, r := make([]int, len(nums)), make([]int, len(nums))
	l[0], r[len(nums)-1] = 1, 1
	for i := 1; i < len(nums); i++ {
		l[i] = l[i-1] * nums[i-1]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		r[i] = r[i+1] * nums[i+1]
	}
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		result[i] = l[i] * r[i]
	}
	return result
}
