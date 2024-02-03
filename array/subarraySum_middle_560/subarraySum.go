package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2023/9/8 下午2:08
 * @Desc:
 */

func main() {
	nums := []int{1}
	k := 0
	fmt.Println(subarraySum(nums, k))
}

func subarraySum(nums []int, k int) int {
	sumMap := make(map[int]int)
	sumMap[0] = 1
	count := 0
	sum := 0
	for _, num := range nums {
		sum += num
		if cnt, ok := sumMap[sum-k]; ok {
			count += cnt
		}
		sumMap[sum]++
	}
	return count
}
