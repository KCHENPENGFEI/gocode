package main

import (
	"fmt"
	"sort"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/8/27 上午9:11
 * @Desc:
 */

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(intervals))
}

func merge(intervals [][]int) [][]int {
	// 按照区间左值排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := make([][]int, 0)
	item := []int{-1, -1}

	for _, interval := range intervals {
		if interval[0] > item[1] {
			// 说明需要新开一个区间
			result = append(result, []int{item[0], item[1]})
			item[0], item[1] = interval[0], interval[1]
		} else {
			// 有重叠需要合并
			if item[1] < interval[1] {
				item[1] = interval[1]
			}
		}
	}
	result = append(result, item)
	return result[1:]
}
