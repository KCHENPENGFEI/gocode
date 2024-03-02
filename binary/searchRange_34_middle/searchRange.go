package main

/**
 * @Author: chenpengfei
 * @Date: 2023/12/17 下午1:42
 * @Desc:给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回 [-1, -1]。

你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。
*/

func main() {

}

func searchRange(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		num := nums[mid]
		if num < target {
			l = mid + 1
		} else if num > target {
			r = mid - 1
		} else {
			// nums分成两个区间分别使用二分查找
			left := findFirstOne(nums[:mid+1], target)
			right := findLastOne(nums[mid:], target)
			return []int{left, right + mid}
		}
	}
	return []int{-1, -1}
}

// 必然存在数值target
func findFirstOne(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2
		num := nums[mid]
		if num < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

// 必然存在数值target
func findLastOne(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l+1)/2
		num := nums[mid]
		if num > target {
			r = mid - 1
		} else {
			l = mid
		}
	}
	return l
}

// searchRange2
//
//	@Description: 二刷，搜索范围优化
//	@param nums
//	@param target
//	@return []int
func searchRange2(nums []int, target int) []int {
	// 找到第一个>=target的位置和最后一个<=target的位置
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	// 第一次二分时候可以记录一下最小的>target的位置
	// 这样第二次二分时候就不需要全量搜索了
	right := len(nums) - 1
	// 第一次二分
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
			if nums[r] > target {
				right = r
			}
		}
	}
	if nums[l] != target {
		return []int{-1, -1}
	}
	var res = []int{l, -1}

	// 第二次二分
	r = right
	for l < r {
		mid := l + (r-l+1)/2
		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid
		}
	}
	res[1] = l
	return res
}
