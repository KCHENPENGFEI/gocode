package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2024/1/28 下午4:01
 * @Desc:
 */

func main() {
	n := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	r := maxSubArray(n)
	fmt.Println(r)
}

func maxSubArray(nums []int) int {
	var imax = make([]int, len(nums))
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			imax[i] = nums[i]
		} else {
			imax[i] = maxTwo(imax[i-1]+nums[i], nums[i])
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
