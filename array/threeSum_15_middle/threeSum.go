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
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
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
