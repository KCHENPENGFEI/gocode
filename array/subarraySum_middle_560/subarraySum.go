package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2023/9/8 下午2:08
 * @Desc: 和为K的子数组
给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。

子数组是数组中元素的连续非空序列。
*/

func main() {
	nums := []int{1}
	k := 0
	fmt.Println(subarraySum(nums, k))
}

// subarraySum
//
//	@Description: 使用前缀和+哈希表解法
//	@param nums
//	@param k
//	@return int
func subarraySum(nums []int, k int) int {
	// map的key为和，value是和的次数
	sumMap := make(map[int]int)
	sumMap[0] = 1
	count := 0
	sum := 0
	for _, num := range nums {
		sum += num
		// sum-k如果出现在map中，说明能有一个子队列和恰好为k
		if cnt, ok := sumMap[sum-k]; ok {
			count += cnt
		}
		sumMap[sum]++
	}
	return count
}
