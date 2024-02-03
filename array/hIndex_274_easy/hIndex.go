package main

import (
	"fmt"
	"sort"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/10/29 下午1:08
 * @Desc:给你一个整数数组 citations ，其中 citations[i] 表示研究者的第 i 篇论文被引用的次数。计算并返回该研究者的 h 指数。
根据维基百科上 h 指数的定义：h 代表“高引用次数” ，一名科研人员的 h 指数 是指他（她）至少发表了 h 篇论文，并且每篇论文 至少 被引用 h 次。如果 h 有多种可能的值，h 指数 是其中最大的那个。
*/

func main() {
	arr := []int{1, 1}
	fmt.Println(hIndex(arr))
}

// hIndex
//
//	@Description: 先进行排序，然后从最后一个元素向前遍历，如果遍历长度<=此时被引用的次数，就将其作为结果
//	@param citations
//	@return int
func hIndex(citations []int) int {
	result := 0
	sort.Ints(citations)
	l := len(citations)
	for i := l - 1; i >= 0; i-- {
		count := citations[i]
		if count >= l-i {
			result = l - i
		}
	}
	return result
}
