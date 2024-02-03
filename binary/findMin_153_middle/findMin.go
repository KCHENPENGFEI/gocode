package main

/**
 * @Author: chenpengfei
 * @Date: 2023/12/17 下午3:53
 * @Desc:
 */

func main() {

}

func findMin(nums []int) int {
	if nums[0] <= nums[len(nums)-1] {
		return nums[0]
	}
	kIndex := findK(nums)
	return nums[kIndex]
}

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
