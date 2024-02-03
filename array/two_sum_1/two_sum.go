package main

import (
	"fmt"
	"sort"
)

/**
 * @Author: chenpengfei
 * @Date: 2022/9/5 10:18 PM
 * @Desc:
 */

func main() {
	//fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	//fmt.Println(twoSum([]int{6, 1, 31, 3}, 9))
	fmt.Println(twoSum2([]int{1, 10, 8, 2, 2, 4}, 11))
}

//
// twoSum
//  @Description:
//  @param nums
//  @param target
//  @return []int
//
func twoSum(nums []int, target int) []int {
	positions := savePosition(nums)
	left, right := findThem(nums, target)
	leftPos := positions[left][0]
	rightPos := positions[right][len(positions[right])-1]
	return []int{leftPos, rightPos}
}

//
// twoSum1
//  @Description: 遍历数组，每次遍历到一个值时，计算与target的差值，在map中查找差值的位置，如果不存在则记录该值的位置到map中
//  @param nums
//  @param target
//  @return []int
//
func twoSum1(nums []int, target int) []int {
	pos := make(map[int]int)
	for i, v := range nums {
		diff := target - v
		if p, ok := pos[diff]; ok {
			return []int{p, i}
		}
		pos[v] = i
	}
	return nil
}

func savePosition(nums []int) map[int][]int {
	var positions = make(map[int][]int)
	for i, v := range nums {
		if _, ok := positions[v]; ok {
			positions[v] = append(positions[v], i)
		} else {
			positions[v] = []int{i}
		}
	}
	return positions
}

func findThem(nums []int, target int) (int, int) {
	sort.Ints(nums)

	start := 0
	end := len(nums) - 1

	for start <= end {
		pre := nums[start]
		after := nums[end]

		sum := pre + after
		if sum == target {
			return pre, after
		}
		if sum < target {
			start += 1
			continue
		}
		if sum > target {
			end -= 1
			continue
		}
	}
	return nums[start], nums[end]
}

//
// twoSum2
//  @Description: 找到这两个值是多少
//  @param nums
//  @param target
//  @return []int
//
func twoSum2(nums []int, target int) []int {
	sort.Ints(nums)
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i]+nums[j] == target {
			return []int{nums[i], nums[j]}
		} else if nums[i]+nums[j] < target {
			i++
		} else {
			j--
		}
	}
	return []int{nums[i], nums[j]}
}
