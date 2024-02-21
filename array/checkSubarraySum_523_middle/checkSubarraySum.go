package main

/**
 * @Author: chenpengfei
 * @Date: 2024/2/17 下午1:09
 * @Desc: 523. 连续的子数组和
给你一个整数数组 nums 和一个整数 k ，编写一个函数来判断该数组是否含有同时满足下述条件的连续子数组：

子数组大小 至少为 2 ，且
子数组元素总和为 k 的倍数。
如果存在，返回 true ；否则，返回 false 。

如果存在一个整数 n ，令整数 x 符合 x = n * k ，则称 x 是 k 的一个倍数。0 始终视为 k 的一个倍数。
*/

// checkSubarraySum
//
//	@Description: 使用一个map来保存余数和位置的映射
//	@param nums
//	@param k
//	@return bool
func checkSubarraySum(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}
	// key为余数
	// value为index下标
	mp := make(map[int]int)
	mp[0] = -1
	sum := 0
	for i, num := range nums {
		sum += num
		r := sum % k
		// 余数相同说明可以被k整除
		// 并且要求i-v长度>=2
		if v, ok := mp[r]; ok {
			if i-v >= 2 {
				return true
			}
		} else {
			// 记录下标
			mp[r] = i
		}
	}
	return false
}
