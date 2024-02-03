package main

import "fmt"

/**
 * @Author: chenpengfei
 * @Date: 2023/12/3 下午2:04
 * @Desc:
 */

func main() {
	a := []int{1, 79, 80, 1, 1, 1, 200, 1}
	k := 3
	fmt.Println(maxScore(a, k))
}

func maxScore(cardPoints []int, k int) int {
	maxS, sum := 0, 0
	length := len(cardPoints)
	for i := 0; i < k; i++ {
		maxS += cardPoints[i]
	}
	sum = maxS
	if length == k {
		return sum
	}

	for i := k - 2; i >= 0; i-- {
		t := (i - (k - 1)) + length
		sum = sum - cardPoints[i+1] + cardPoints[t]
		if sum > maxS {
			maxS = sum
		}
	}

	sum = sum - cardPoints[0] + cardPoints[length-k]
	if sum > maxS {
		maxS = sum
	}
	return maxS
}
