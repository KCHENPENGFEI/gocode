package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2022/9/5 10:40 PM
 * @Desc:
 */

func main() {
	nums := []int{2, 3, 3, 2, 7, 3, 4}
	val := 3
	length := removeElement(nums, val)
	fmt.Println(length)
	fmt.Println(nums[0:length])

}

//
// removeElement
//  @Description: 使用双指针，前指针h从头开始遍历，后指针t永远指向数组中最后一个不等于val的值，
//  如果nums[h] == val，则将nums[h]和nums[t]互换位置，然后重新计算t的位置，最后返回t+1即可
//  @param nums
//  @param val
//  @return int
//
func removeElement(nums []int, val int) int {
	start := 0
	end := findLastIndexNotEqualVal(len(nums)-1, nums, val)
	if end < 0 {
		return 0
	}
	if end == start {
		return start + 1
	}
	for start < end {
		if nums[start] == val {
			// 与end位置互换
			nums[start], nums[end] = nums[end], nums[start]
			start++
			end = findLastIndexNotEqualVal(end-1, nums, val)
		} else {
			start++
		}
	}
	return end + 1
}

//
// findLastIndexNotEqualVal
//  @Description: 找到最后一个不等于val值的位置
//  @param start 起始位置(从后往前)
//  @param nums
//  @param val
//  @return int
//
func findLastIndexNotEqualVal(start int, nums []int, val int) int {
	for i := start; i >= 0; i-- {
		if nums[i] != val {
			return i
		}
	}
	// 没有找到不等于val的值
	return -1
}
