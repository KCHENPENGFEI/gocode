package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2024/1/27 下午5:02
 * @Desc:
 */

func main() {
	s := "eccbbbbdec"
	r := partitionLabels(s)
	fmt.Println(r)
}

func partitionLabels(s string) []int {
	// 整理成区间数组
	intervals := buildIntervals(s)
	var partition = make([][]int, 0)
	item := []int{-1, -1}
	for _, inter := range intervals {
		if inter[0] > item[1] {
			// 新开区间
			partition = append(partition, []int{item[0], item[1]})
			item[0], item[1] = inter[0], inter[1]
		} else {
			// 重叠合并
			if item[1] < inter[1] {
				item[1] = inter[1]
			}
		}
	}
	partition = append(partition, item)
	partition = partition[1:]

	var res []int
	for _, inter := range partition {
		res = append(res, inter[1]-inter[0]+1)
	}
	return res
}

func buildIntervals(s string) [][]int {
	var res [][]int
	posMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if pos, ok := posMap[s[i]]; ok {
			res[pos][1] = i
		} else {
			res = append(res, []int{i, i})
			posMap[s[i]] = len(res) - 1
		}
	}
	return res
}
