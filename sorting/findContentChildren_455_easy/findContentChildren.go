package main

import (
	"fmt"
	"sort"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/8/5 下午1:36
 * @Desc:
 */

func main() {
	g := []int{1, 2, 3}
	s := []int{1, 2, 3, 4}
	fmt.Println(findContentChildren(g, s))
}

func findContentChildren(g []int, s []int) int {
	// 先排序，后贪心
	sort.Ints(g)
	sort.Ints(s)

	j := 0
	for i := 0; i < len(g); i++ {
		if j >= len(s) {
			// 没有饼干了
			return i
		}

		if s[j] >= g[i] {
			// 能满足
			j++
			continue
		} else {
			// 孩子还没分配，所以i--
			j++
			i--
		}
	}
	return len(g)
}
