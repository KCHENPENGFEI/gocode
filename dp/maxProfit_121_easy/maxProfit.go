package main

/**
 * @Author: chenpengfei
 * @Date: 2024/1/27 下午3:34
 * @Desc:
 */

// maxProfit
//
//	@Description: 贪心算法循环一次
//	@param prices
//	@return int
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	res := 0
	minVal := prices[0]
	for i := 1; i < len(prices); i++ {
		res = maxTwo(res, prices[i]-minVal)
		minVal = minTwo(minVal, prices[i])
	}
	return res
}

func maxTwo(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minTwo(a, b int) int {
	if a < b {
		return a
	}
	return b
}
