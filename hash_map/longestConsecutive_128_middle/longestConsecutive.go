package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2023/8/27 上午10:34
 * @Desc:
 */

func main() {
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Println(longestConsecutive(nums))
}

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
