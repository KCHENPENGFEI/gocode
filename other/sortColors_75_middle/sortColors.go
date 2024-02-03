package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2024/2/3 下午6:56
 * @Desc: 颜色分类
给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

必须在不使用库内置的 sort 函数的情况下解决这个问题。
*/

func main() {
	nums := []int{2, 0, 1}
	sortColors(nums)
	fmt.Println(nums)
}

func sortColors(nums []int) {
	// 遇到2就向后置换
	// 遇到0就向前置换
	backIdx := len(nums) - 1
	headIdx := 0
	// i<=backIndex可以减少循环次数
	for i := 0; i <= backIdx; i++ {
		if nums[i] == 2 {
			nums[i], nums[backIdx] = nums[backIdx], nums[i]
			backIdx--
			// 这里需要进行i--，向后置换的数据是没有遍历过的
			i--
		} else if nums[i] == 0 {
			nums[i], nums[headIdx] = nums[headIdx], nums[i]
			// 这里不需要进行i--，因为向前置换的值都是遍历过的
			headIdx++
		}
	}
}
