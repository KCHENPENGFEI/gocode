package main

import (
	"fmt"
	"sort"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/3/10 下午11:17
 * @Desc: 整数数组的一个 排列  就是将其所有成员以序列或线性顺序排列。

例如，arr = [1,2,3] ，以下这些都可以视作 arr 的排列：[1,2,3]、[1,3,2]、[3,1,2]、[2,3,1] 。
整数数组的 下一个排列 是指其整数的下一个字典序更大的排列。更正式地，如果数组的所有排列根据其字典顺序从小到大排列在一个容器中，那么数组的 下一个排列 就是在这个有序容器中排在它后面的那个排列。如果不存在下一个更大的排列，那么这个数组必须重排为字典序最小的排列（即，其元素按升序排列）。

例如，arr = [1,2,3] 的下一个排列是 [1,3,2] 。
类似地，arr = [2,3,1] 的下一个排列是 [3,1,2] 。
而 arr = [3,2,1] 的下一个排列是 [1,2,3] ，因为 [3,2,1] 不存在一个字典序更大的排列。
给你一个整数数组 nums ，找出 nums 的下一个排列。

必须 原地 修改，只允许使用额外常数空间。

*/

func main() {
	nums := []int{1, 2}
	nextPermutation(nums)
	fmt.Println(nums)
}

func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}
	for i := len(nums) - 1; i >= 1; i-- {
		// 找到降序的index
		if nums[i-1] < nums[i] {
			helper(i-1, nums)
			return
		}
	}
	// 全降序则重新排序即可
	sort.Ints(nums)
}

//
// helper
//  @Description: 这里要注意的是，找到index后并不是直接和index+1位置的值交换
//  @param index
//  @param nums
//
func helper(index int, nums []int) {
	minVal := nums[index+1]
	minValIndex := index + 1
	for i := index + 1; i < len(nums); i++ {
		if nums[i] > nums[index] && nums[i] < minVal {
			minVal = nums[i]
			minValIndex = i
		}
	}
	// 交换位置
	nums[index], nums[minValIndex] = nums[minValIndex], nums[index]
	// 排序
	sort.Ints(nums[index+1:])
}
