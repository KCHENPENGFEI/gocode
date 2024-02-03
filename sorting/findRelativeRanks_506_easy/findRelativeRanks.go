package main

import (
	"fmt"
	"sort"
	"strconv"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/8/5 下午1:58
 * @Desc: 给你一个长度为 n 的整数数组 score ，其中 score[i] 是第 i 位运动员在比赛中的得分。所有得分都 互不相同 。
	运动员将根据得分 决定名次 ，其中名次第 1 的运动员得分最高，名次第 2 的运动员得分第 2 高，依此类推。
*/

func main() {
	score := []int{5, 6, 3, 2, 1}
	fmt.Println(findRelativeRanks(score))
}

func findRelativeRanks(score []int) []string {
	var ss = make([]int, len(score))
	// 复制一份数据
	copy(ss, score)
	// 按照分数倒排
	sort.Slice(score, func(i, j int) bool {
		return score[i] > score[j]
	})

	// 计算分数排名
	rank := make(map[int]int, len(score))
	for i, s := range score {
		rank[s] = i
	}

	var result = make([]string, len(score))
	// 遍历运动员成绩得到排名
	for i, s := range ss {
		iRank := rank[s]

		if iRank == 0 {
			result[i] = "Gold Medal"
		} else if iRank == 1 {
			result[i] = "Silver Medal"
		} else if iRank == 2 {
			result[i] = "Bronze Medal"
		} else {
			result[i] = strconv.Itoa(iRank + 1)
		}
	}
	return result
}
