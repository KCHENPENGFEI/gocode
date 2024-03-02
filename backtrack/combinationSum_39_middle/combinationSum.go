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

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var path []int
	result := new([][]int)
	helper(candidates, path, target, 0, 0, result)
	return *result
}

// helper
//
//	@Description: 二刷，为了防止重复答案生成因此要排序，并且保证每次递归index递增
//	@param nums
//	@param path
//	@param target
//	@param sum
//	@param start
//	@param res
func helper(nums []int, path []int, target, sum, start int, res *[][]int) {
	if target == sum {
		c := make([]int, len(path))
		copy(c, path)
		*res = append(*res, c)
		return
	}

	for i := start; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			// 重复不再处理
			continue
		}
		if sum+nums[i] > target {
			// 剪枝
			return
		} else {
			sum += nums[i]
			path = append(path, nums[i])
			// start递增
			helper(nums, path, target, sum, i, res)
		}
		// 回退
		sum -= nums[i]
		path = path[:len(path)-1]
	}
}
