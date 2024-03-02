package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2023/12/17 下午4:25
 * @Desc:
 */

func main() {
	nums := []int{5, 4, 6, 2}
	r := permute(nums)
	//for _, d := range r {
	//	fmt.Printf("%p: \n", d)
	//}
	fmt.Println(r)
}

func permute(nums []int) [][]int {
	var track []int
	var trackSet = make(map[int]struct{})
	var n = len(nums)
	var result [][]int
	dfs(track, trackSet, nums, n, &result)
	return result
}

func dfs(track []int, trackSet map[int]struct{}, nums []int, n int, result *[][]int) {
	if len(track) == n {
		// 需要对track进行copy，因为track底层是指针类型
		var cp = make([]int, len(track))
		copy(cp, track)
		*result = append(*result, cp)
		return
	}
	for i := 0; i < n; i++ {
		num := nums[i]
		if _, ok := trackSet[num]; ok {
			continue
		}
		trackSet[num] = struct{}{}
		track = append(track, num)
		dfs(track, trackSet, nums, n, result)
		track = track[:len(track)-1]
		delete(trackSet, num)
	}
}

// permute2
//
//	@Description: 二刷
//	@param nums
//	@return [][]int
func permute2(nums []int) [][]int {
	set := make(map[int]struct{})
	var path []int
	result := new([][]int)
	trackback(nums, set, path, result)
	return *result
}

func trackback(nums []int, set map[int]struct{}, path []int, res *[][]int) {
	if len(path) == len(nums) {
		c := make([]int, len(path))
		copy(c, path)
		*res = append(*res, c)
		return
	}

	for i := 0; i < len(nums); i++ {
		if _, ok := set[nums[i]]; !ok {
			set[nums[i]] = struct{}{}
			path = append(path, nums[i])
			trackback(nums, set, path, res)
			path = path[:len(path)-1]
			delete(set, nums[i])
		}
	}
}
