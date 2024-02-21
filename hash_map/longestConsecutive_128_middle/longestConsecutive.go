package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2023/8/27 上午10:34
 * @Desc: 最长连续序列
 */

func main() {
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Println(longestConsecutive(nums))
}

// longestConsecutive
//
//		@Description: 做法: 1. 首先使用HashSet对数组进行去重
//	  - 2. 遍历数组每一个元素i，如果HashSet中存在i-1这个数字，说明元素i肯定不是构成最长序列的第一个元素，直接跳过
//	  - 3. 如果元素i是第一个元素，那么考察i+1、i+2、...是否在HashSet中，然后统计长度即可
//	    @param nums
//	    @return int
func longestConsecutive(nums []int) int {
	hashMap := make(map[int]struct{})
	maxLen := 0
	for _, num := range nums {
		hashMap[num] = struct{}{}
	}

	for _, num := range nums {
		if _, ok := hashMap[num-1]; ok {
			continue
		}
		length := 1
		for {
			if _, ok := hashMap[num+1]; ok {
				num++
				length++
			} else {
				break
			}
		}
		if length > maxLen {
			maxLen = length
		}
	}

	return maxLen
}

func longestConsecutive1(nums []int) int {
	hashMap := make(map[int]bool)
	maxLen := 0
	for _, num := range nums {
		hashMap[num] = true
	}

	for _, num := range nums {
		if _, ok := hashMap[num-1]; ok {
			continue
		}
		length := 1
		for {
			if _, ok := hashMap[num+1]; ok {
				num++
				length++
			} else {
				break
			}
		}
		if length > maxLen {
			maxLen = length
		}
	}

	return maxLen
}
