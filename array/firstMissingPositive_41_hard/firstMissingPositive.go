package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2023/10/15 下午2:21
 * @Desc: 给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。

请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
原地哈希，假设一个数字是d，那么把他移动到下标d-1的位置上
如果d<=0或者d>n就说明没有位置也不可能作为答案
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
			// 位置i处的数值已经是正确的
			// 如果小于0则不修改，意为该位置还没有正确的数字
			i++
			continue
		}
		if nums[i] > n {
			// 说明nums[i]已经超出n的范围，nums[i]不可能是答案
			// 该位置也可以清空
			nums[i] = 0
		} else {
			// 找到nums[i]应该在的下标的位置
			index := nums[i] - 1
			if nums[index] == index+1 {
				// 说明nums[i]这个数值有重复，也可以清空这个位置
				nums[i] = 0
			}
			// 将nums[i]移动到他应该在的位置上
			nums[i], nums[index] = nums[index], nums[i]
			// 这里不能进行i++
			// 因为刚才把i位置上数字移走了，但是又移过来一个新的数字，仍然需要处理这个数字
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			return i + 1
		}
	}
	return n + 1
}
