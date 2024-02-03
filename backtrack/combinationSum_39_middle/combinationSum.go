package main

import (
	"fmt"
	"sort"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/12/30 下午2:05
 * @Desc:
 */

func main() {
	c := []int{2, 3, 5}
	t := 7
	r := combinationSum(c, t)
	fmt.Println(r)
}

func combinationSum(candidates []int, target int) [][]int {
	// 先做排序
	sort.Ints(candidates)
	var track []int
	var result [][]int
	dfs(track, candidates, 0, target, &result)
	return result
}

func dfs(track, candidates []int, start, target int, result *[][]int) {
	if target == 0 {
		// 找到一组结果
		var cp = make([]int, len(track))
		copy(cp, track)
		*result = append(*result, cp)
		return
	}

	// 引入start确保答案都是升序，保证答案不重复
	for i := start; i < len(candidates); i++ {
		track = append(track, candidates[i])
		target = target - candidates[i]
		if target < 0 {
			// 直接剪枝
			break
		} else {
			dfs(track, candidates, i, target, result)
		}
		track = track[:len(track)-1]
		target = target + candidates[i]
	}
}
