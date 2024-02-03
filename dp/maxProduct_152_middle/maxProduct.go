package main

/**
 * @Author: chenpengfei
 * @Date: 2024/1/28 下午3:47
 * @Desc:
 */

func main() {

}

func maxProduct(nums []int) int {
	// 定义imax[i]是包含了nums[i]情况下最大乘积
	// 定义imin[i]是包含了nums[i]情况下最小乘积
	var imax = make([]int, len(nums))
	var imin = make([]int, len(nums))
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			imax[i], imin[i] = nums[i], nums[i]
		} else {
			if nums[i] < 0 {
				imax[i] = maxTwo(imin[i-1]*nums[i], nums[i])
				imin[i] = minTwo(imax[i-1]*nums[i], nums[i])
			} else {
				imax[i] = maxTwo(imax[i-1]*nums[i], nums[i])
				imin[i] = minTwo(imin[i-1]*nums[i], nums[i])
			}
		}
		res = maxTwo(res, imax[i])
	}
	return res
}

func maxTwo(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minTwo(a, b int) int {
	if a < b {
		return a
	}
	return b
}
