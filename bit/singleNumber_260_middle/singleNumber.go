package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2023/10/14 下午4:45
 * @Desc:
 */

func main() {
	nums := []int{1, 2, 1, 3, 2, 5, 3, 5, 7, 9}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) []int {
	x := 0
	for _, num := range nums {
		x = x ^ num
	}
	t := x & -x
	a1, a2 := 0, 0
	for _, num := range nums {
		if num&t == t {
			a1 = a1 ^ num
		} else {
			a2 = a2 ^ num
		}
	}
	return []int{a1, a2}
}
