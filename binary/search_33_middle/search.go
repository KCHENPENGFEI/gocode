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

func search(nums []int, target int) int {
	// 要做两次判断
	// 一次判断当前mid在左半段还是右半段内
	// 一次判断当前mid和target之间关系
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		// 判断在哪半段
		if nums[mid] >= nums[l] {
			// 左半段
			if target > nums[mid] || target < nums[l] {
				// 有两种情况需要l右移
				l = mid + 1
			} else {
				r = mid
			}
		} else {
			// 右半段
			if target > nums[mid] && target < nums[l] {
				// 有两种情况需要l右移
				l = mid + 1
			} else {
				r = mid
			}
		}
	}
	if nums[l] == target {
		return l
	}
	return -1
}
