package main

/**
 * @Author: chenpengfei
 * @Date: 2024/1/28 下午3:00
 * @Desc:
 */

func main() {

}

func coinChange(coins []int, amount int) int {
	var dp = make([]int, amount+1)
	for i := 1; i < amount+1; i++ {
		dp[i] = amount + 1
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = minTwo(dp[i], dp[i-coin]+1)
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func minTwo(a, b int) int {
	if a < b {
		return a
	}
	return b
}
