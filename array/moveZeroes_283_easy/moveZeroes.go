package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2023/9/3 下午3:11
 * @Desc: 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
	请注意 ，必须在不复制数组的情况下原地对数组进行操作。
*/

func main() {
	nums := []int{0, 2}
	moveZeroes(nums)
	fmt.Println(nums)
}

// 双指针i，j，其中i指向第一个为0的位置，j指向i之后第一个不为0的位置，然后交换i和j的内容即可
func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	i, j := 0, 0
	for j < len(nums) {
		// 找到第一个为0的位置
		i = findI(nums, i)
		if i == -1 {
			return
		}
		// 找到i之后的第一个不为0的位置
		j = findJ(nums, i+1)
		if j == -1 {
			return
		}

		// 交换i和j的位置
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func findI(nums []int, start int) int {
	for i := start; i < len(nums); i++ {
		if nums[i] == 0 {
			return i
		}
	}
	return -1
}

func findJ(nums []int, start int) int {
	for i := start; i < len(nums); i++ {
		if nums[i] != 0 {
			return i
		}
	}
	return -1
}
