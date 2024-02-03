package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2022/9/7 11:15 PM
 * @Desc:
 */

func main() {
	nums := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump1(nums))
}

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	return withIndexJump(nums, 0)
}

//
// canJump1
//  @Description: 维持一个最远距离farthest，遍历数组不断更新这个值，如果位置i>farthest，说明不可达，返回false
//  @param nums
//  @return bool
//
func canJump1(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	farthest := 0
	for i := 0; i < len(nums); i++ {
		if i > farthest {
			// 位置i不可达
			return false
		}
		farthest = max(farthest, i+nums[i])
	}
	return true
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//
// withIndexJump
//  @Description: 递归写法，会有很多的重复计算，可以使用数组将计算结果保存起来避免继续递归
//  @param nums
//  @param index
//  @return bool
//
func withIndexJump(nums []int, index int) bool {
	length := len(nums)
	num := nums[index]
	if (length-1)-index <= num {
		// 可以跳到最后
		return true
	}
	// 继续跳
	for i := num; i > 0; i-- {
		ret := withIndexJump(nums, index+i)
		if ret {
			return true
		}
	}
	// 跳完都找不到
	return false
}
