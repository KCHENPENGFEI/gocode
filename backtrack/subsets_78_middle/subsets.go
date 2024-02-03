package main

import (
	"fmt"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/12/17 下午5:20
 * @Desc:
 */

func main() {
	nums := []int{1, 2, 3, 4}
	r := subsets(nums)
	fmt.Println(r)
}

func subsets(nums []int) [][]int {
	var (
		track  []int
		result [][]int
	)
	for i := 0; i <= len(nums); i++ {
		dfs(track, nums, 0, i, &result)
	}
	return result
}

func dfs(track []int, nums []int, start, n int, res *[][]int) {
	if len(track) == n {
		var cp = make([]int, len(track))
		copy(cp, track)
		*res = append(*res, cp)
		return
	}
	for i := start; i < len(nums); i++ {
		track = append(track, nums[i])
		dfs(track, nums, i+1, n, res)
		track = track[:len(track)-1]
	}
}
