package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2023/8/19 下午3:21
 * @Desc:
 */

func main() {
	s := "DDDDD"
	fmt.Println(diStringMatch(s))
}

func diStringMatch(s string) []int {
	var result []int
	max, min := len(s), 0
	for _, c := range s {
		if c == 'D' {
			result = append(result, max)
			max--
		} else {
			result = append(result, min)
			min++
		}
	}
	result = append(result, max)
	return result
}
