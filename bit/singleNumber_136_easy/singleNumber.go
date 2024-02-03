package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2023/10/14 下午4:02
 * @Desc:
 */

func main() {
	nums := []int{10, 10, 1, 3, 3, 4, 5, 5, 4}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result = result ^ num
	}
	return result
}
