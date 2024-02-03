package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2023/3/14 下午10:42
 * @Desc: 整数数组 nums 按升序排列，数组中的值 互不相同 。

在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。

给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。

你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。

*/

func main() {
	nums := []int{1, 3}
	target := 0
	fmt.Println(search(nums, target))
}

// search
//
//	@Description: 二分法，难点在于判断何时向左分，何时向右分
//	@param nums
//	@param target
//	@return int
func search(nums []int, target int) int {
	if len(nums) == 1 {
		if target == nums[0] {
			return 0
		}
		return -1
	}

	head := nums[0]
	tail := nums[len(nums)-1]
	if head < tail {
		return searchInOrder(nums, target, 0, len(nums)-1)
	}
	if target == head {
		return 0
	} else if target < head {
		// 在右半边
		l := 0
		r := len(nums) - 1
		for l < r {
			m := (l + r) / 2
			if nums[m] == target {
				return m
			}
			if nums[m] > target && nums[m] > nums[r] {
				// 还在右半边
				l = m + 1
			} else if nums[m] > target && nums[m] < nums[r] {
				r = m
			} else if nums[m] < target && nums[m] > nums[r] {

			} else {
				return searchInOrder(nums, target, m+1, r)
			}
		}
		if nums[r] == target {
			return r
		}
		return -1
	} else {
		// 在左半边
		l := 0
		r := len(nums) - 1
		for l < r {
			m := (l + r) / 2
			if nums[m] == target {
				return m
			}
			if nums[m] > target && nums[l] < nums[m] {
				return searchInOrder(nums, target, l, m)
			} else if nums[m] > target && nums[l] > nums[m] {

			} else if nums[m] < target && nums[l] < nums[m] {
				l = m + 1
			} else {
				r = m
			}
		}
		if nums[r] == target {
			return r
		}
		return -1
	}
}

func searchInOrder(nums []int, target int, start, end int) int {
	for start < end {
		m := (start + end) / 2
		if nums[m] == target {
			return m
		}
		if nums[m] < target {
			start = m + 1
		} else {
			end = m
		}
	}
	if nums[end] == target {
		return end
	}
	return -1
}

func search1(nums []int, target int) int {
	if nums[0] <= nums[len(nums)-1] {
		return searchInOrder1(nums, target)
	}
	kIndex := findK(nums)
	fmt.Println(kIndex)
	if target <= nums[len(nums)-1] {
		// 右边
		index := searchInOrder1(nums[kIndex:], target)
		if index == -1 {
			return -1
		}
		return index + kIndex
	} else {
		return searchInOrder1(nums[:kIndex], target)
	}
}

// 找到最小值的坐标
func findK(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] < nums[r] {
			// 左边
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func searchInOrder1(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}
