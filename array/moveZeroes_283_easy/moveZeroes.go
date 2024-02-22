package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2023/9/3 下午3:11
 * @Desc: 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
	请注意 ，必须在不复制数组的情况下原地对数组进行操作。
*/

func main() {
	nums := []int{0, 1, 9, 3, 12}
	moveZeroes2(nums)
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

// moveZeroes2
//
//	@Description: 二刷
//	@param nums
func moveZeroes2(nums []int) {
	fristZero := -1
	for i := 0; i < len(nums); i++ {
		if fristZero == -1 && nums[i] == 0 {
			fristZero = i
		}
		if nums[i] != 0 && fristZero != -1 {
			// 交换位置
			nums[i], nums[fristZero] = nums[fristZero], nums[i]
			// 因为交换了位置所以当前fristZero肯定不是0了，++之后必然是0的位置
			// 如果交换的位置相邻，那么++之后肯定是0
			// 如果交换的位置不相邻，那么++之后就是没有移动过的数组的0
			fristZero++
		}
	}
}
