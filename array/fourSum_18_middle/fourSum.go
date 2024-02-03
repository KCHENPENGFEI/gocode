package main

import (
	"fmt"
	"sort"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/3/10 下午10:05
 * @Desc:
 */

func main() {
	fmt.Println(fourSum([]int{0, 0, 0, 0, 0, 0, 0, 0, 0}, 0))
}

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}
	length := len(nums)
	// 排序
	sort.Ints(nums)
	var result [][]int
	for first := 0; first < length-3; first++ {
		if (first != 0) && (nums[first] == nums[first-1]) {
			continue
		}
		for second := first + 1; second < length-2; second++ {
			if (second != first+1) && (nums[second] == nums[second-1]) {
				continue
			}
			for third := second + 1; third < length-1; third++ {
				if (third != second+1) && (nums[third] == nums[third-1]) {
					continue
				}
				fourth := length - 1
				for fourth > third {
					if nums[first]+nums[second]+nums[third]+nums[fourth] > target {
						fourth--
					} else if nums[first]+nums[second]+nums[third]+nums[fourth] == target {
						result = append(result, []int{nums[first], nums[second], nums[third], nums[fourth]})
						break
					} else {
						break
					}
				}
			}
		}
	}
	return result
}
