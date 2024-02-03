package main

import (
	"fmt"
	"sort"
)

/**
 * @Author: chenpengfei
 * @Date: 2023/8/20 下午12:32
 * @Desc:
 */

func main() {
	points := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	fmt.Println(findMinArrowShots1(points))
}

func findMinArrowShots1(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool { return points[i][1] < points[j][1] })
	maxRight := points[0][1]
	ans := 1
	for _, p := range points {
		if p[0] > maxRight {
			maxRight = p[1]
			ans++
		}
	}
	return ans
}

func findMinArrowShots(points [][]int) int {
	// 按照右边界进行升序排列
	sort.Slice(points, func(i, j int) bool {
		iLeft, yLeft := points[i][0], points[j][0]
		iRight, yRight := points[i][1], points[j][1]
		if iRight < yRight {
			return true
		} else if iRight > yRight {
			return false
		} else {
			return iLeft < yLeft
		}
	})

	var deleted = make([]bool, len(points))
	result := 0
	// 不断寻找所有右边界的最左值
	for {
		index := 0
		found := false
		for index = 0; index < len(points); index++ {
			if deleted[index] == true {
				continue
			}
			result++
			found = true
			break
		}
		if !found {
			return result
		}
		for j := index; j < len(points); j++ {
			right := points[index][1]
			if right <= points[j][1] && right >= points[j][0] {
				deleted[j] = true
			}
		}
	}
}
