package main

import (
	"fmt"
	"sort"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/9/3 下午3:45
 * @Desc:
 */

func main() {
	nums := []int{}
	fmt.Println(threeSum2(nums))
}

type triple struct {
	a int
	b int
	c int
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	m := make(map[int][]int)
	for i, num := range nums {
		if m[num] == nil {
			m[num] = []int{i}
		} else {
			m[num] = append(m[num], i)
		}
	}

	var r = make([][]int, 0)
	set := make(map[triple]bool)

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			return r
		}
		if i != 0 && (nums[i] == nums[i-1] && nums[i] == nums[i+1]) {
			continue
		}
		for j := i + 1; j < len(nums)-1; j++ {
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}
			target := 0 - nums[i] - nums[j]
			if m[target] == nil {
				continue
			}
			for _, index := range m[target] {
				t := triple{a: nums[i], b: nums[j], c: target}
				if index > j && set[t] == false {
					r = append(r, []int{nums[i], nums[j], target})
					set[t] = true
				}
			}
		}
	}
	return r

}

// threeSum2
//
//	@Description: 二刷，双指针
//	@param nums
//	@return [][]int
func threeSum2(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			// 剪枝快速返回
			return result
		}
		if i > 0 && nums[i] == nums[i-1] {
			// 重复
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				// 说明找到了有效的三元组
				result = append(result, []int{nums[i], nums[j], nums[k]})
				for j < k && nums[j] == nums[j+1] {
					// j指针重复则继续累加，找到第一个不重复的
					j++
				}
				for j < k && nums[k] == nums[k-1] {
					// 同理找到第一个不重复的k
					k--
				}
				// 在这里才找到
				j++
				k--
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return result

}
